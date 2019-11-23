package api_test

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/smarkm/k8s-nms/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/yaml"
)

var _ = Describe("Network", func() {
	var (
		clientset *kubernetes.Clientset
	)
	BeforeEach(func() {
		kconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err := clientcmd.BuildConfigFromFlags("", kconfig)
		Expect(err).NotTo(HaveOccurred())
		clientset, err = kubernetes.NewForConfig(config)
	})
	It("Init Client set", func() {
		Expect(clientset).NotTo(Equal(nil))
		api.PodNetwork(clientset)
	})
	It("List Nodes", func() {
		nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
		Expect(err).NotTo(HaveOccurred())
		fmt.Printf("Nodes: %d\n", len(nodes.Items))

	})
	It("Pod network", func() {
		cm, err := clientset.CoreV1().ConfigMaps("kube-system").Get("kubeadm-config", metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		clc := cm.Data["ClusterConfiguration"]
		t := struct {
			Networking struct {
				PodSubnet     string
				ServiceSubnet string
			}
			ApiVersion string
		}{}
		err = yaml.Unmarshal([]byte(clc), &t)
		Expect(err).NotTo(HaveOccurred())
		fmt.Println("Pod Network and service network", t)
	})

})
