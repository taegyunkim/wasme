package cache_test

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/solo-io/wasme/pkg/consts/test"
	testutils "github.com/solo-io/wasme/test"

	corev1 "k8s.io/api/core/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/go-utils/randutils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"

	. "github.com/solo-io/wasme/pkg/cache"
)

var _ = Describe("Deploy", func() {
	var (
		kube kubernetes.Interface
		// switch to running the test vs a real kube cluster
		useRealKube = os.Getenv("USE_REAL_KUBE") != ""

		cacheNamespace = "wasme-cache-test-" + randutils.RandString(4)

		testImage = test.IstioAssemblyScriptImage

		// used for pushing images when USE_REAL_KUBE is true
		operatorImage = func() string {
			if gcloudProject := os.Getenv("GCLOUD_PROJECT_ID"); gcloudProject != "" {
				return fmt.Sprintf("gcr.io/%v/wasme", gcloudProject)
			}
			return "quay.io/solo-io/wasme"
		}()
	)

	BeforeEach(func() {
		if useRealKube {
			err := testutils.RunMake("wasme-image", func(cmd *exec.Cmd) {
				cmd.Args = append(cmd.Args, "OPERATOR_IMAGE="+operatorImage)
				cmd.Args = append(cmd.Args, "VERSION="+cacheNamespace)
			})
			Expect(err).NotTo(HaveOccurred())

			err = testutils.RunMake("wasme-image-push", func(cmd *exec.Cmd) {
				cmd.Args = append(cmd.Args, "OPERATOR_IMAGE="+operatorImage)
				cmd.Args = append(cmd.Args, "VERSION="+cacheNamespace)
			})
			Expect(err).NotTo(HaveOccurred())

			cfg, err := kubeutils.GetConfig("", "")
			Expect(err).NotTo(HaveOccurred())

			kube, err = kubernetes.NewForConfig(cfg)
			Expect(err).NotTo(HaveOccurred())
		} else {
			kube = fake.NewSimpleClientset()
		}
	})
	AfterEach(func() {
		kube.AppsV1().DaemonSets(cacheNamespace).Delete(CacheName, nil)
		kube.CoreV1().ConfigMaps(cacheNamespace).Delete(CacheName, nil)
		kube.CoreV1().Namespaces().Delete(cacheNamespace, nil)
	})
	It("creates the cache namespace, configmap, and daemonset", func() {

		deployer := NewDeployer(kube, cacheNamespace, "", operatorImage, cacheNamespace, nil, corev1.PullAlways)

		err := deployer.EnsureCache()
		Expect(err).NotTo(HaveOccurred())

		_, err = kube.CoreV1().Namespaces().Get(cacheNamespace, v1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())

		cm, err := kube.CoreV1().ConfigMaps(cacheNamespace).Get(CacheName, v1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())

		_, err = kube.AppsV1().DaemonSets(cacheNamespace).Get(CacheName, v1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())

		if !useRealKube {
			return
		}

		// multiple runs should not error
		err = deployer.EnsureCache()
		Expect(err).NotTo(HaveOccurred())

		// eventually pods should be ready
		Eventually(func() (int32, error) {
			cacheDaemonSet, err := kube.AppsV1().DaemonSets(cacheNamespace).Get(CacheName, v1.GetOptions{})
			if err != nil {
				return 0, err
			}
			return cacheDaemonSet.Status.NumberReady, nil
		}, time.Second*30).Should(Equal(int32(1)))

		// test that cache event is fired after updating the config
		cm.Data[ImagesKey] = testImage
		_, err = kube.CoreV1().ConfigMaps(cacheNamespace).Update(cm)
		Expect(err).NotTo(HaveOccurred())

		var events []corev1.Event
		// eventually event should be fired with success
		Eventually(func() ([]corev1.Event, error) {
			events, err = GetImageEvents(kube, cacheNamespace, testImage)
			return events, err
		}, time.Second*30).Should(HaveLen(1))

		Expect(events[0].Reason).To(Equal(Reason_ImageAdded))
	})
})
