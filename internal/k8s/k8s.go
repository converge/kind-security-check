package k8s

import (
	"context"
	"errors"
	"github.com/converge/kind-security-check/pkg"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"net"
)

type KubernetesInterface interface {
	CoreV1() v1.CoreV1Interface
}

type KubernetesClient struct {
	clientSet KubernetesInterface
}

// NewKubernetesClient creates a new instance of KubernetesClient.
// It takes an input of type KubernetesInterface and returns an instance of KubernetesClient.
func NewKubernetesClient(clientset KubernetesInterface) (KubernetesClient, error) {

	k8sClient := KubernetesClient{
		clientSet: clientset,
	}
	return k8sClient, nil
}

// CheckPodsInDefaultNamespace checks if there are deployed pods in the default namespace.
func (k8sClient *KubernetesClient) CheckPodsInDefaultNamespace() error {

	ctx := context.Background()
	pods, err := k8sClient.clientSet.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, item := range pods.Items {
		if item.Namespace == "default" {
			return errors.New("there should not have pods in default namespace [link reference #100]")
		}
	}

	return nil

}

// CheckExposeControlPlane checks if the Kubernetes control plane is exposed to public IP addresses.
func (k8sClient *KubernetesClient) CheckExposeControlPlane() error {

	nodes, err := k8sClient.clientSet.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, node := range nodes.Items {
		if node.ObjectMeta.Labels["kubernetes.io/hostname"] == "kind-control-plane" {

			for _, address := range node.Status.Addresses {
				var ipAddress net.IP
				if ipAddress = net.ParseIP(address.Address); ipAddress == nil {
					continue
				}

				if pkg.IsPublicIP(ipAddress) {
					return errors.New(
						"control plane should not be exposed with a public IP address [link reference #101]",
					)
				}
			}
		}

	}
	return nil
}
