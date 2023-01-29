package k8s

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	_ "k8s.io/client-go/kubernetes/fake"
	"reflect"
	"testing"
)

func TestNewKubernetesClient(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	k8sClient, err := NewKubernetesClient(clientSet)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name      string
		clientSet KubernetesInterface
		want      KubernetesClient
		wantErr   bool
	}{
		{
			name:      "test new k8s client",
			clientSet: clientSet,
			want:      k8sClient,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKubernetesClient(tt.clientSet)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKubernetesClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKubernetesClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKubernetesClient_CheckPodsInDefaultNamespace(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	k8sClient, err := NewKubernetesClient(clientSet)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name    string
		pod     *corev1.Pod
		wantErr bool
	}{
		{
			name:    "pass if there is NO pod running on default namespace",
			wantErr: false,
			pod:     &corev1.Pod{},
		},
		{
			name:    "fail if there is some pod running on default namespace",
			wantErr: true,
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-pod",
					Labels: map[string]string{
						"app": "test",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test-container",
							Image: "test-image",
						},
					},
				},
				Status: corev1.PodStatus{
					Phase: corev1.PodRunning,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// if pod testing pod has running state, create it
			if tt.pod.Status.Phase == corev1.PodRunning {
				_, err = clientSet.CoreV1().Pods("default").Create(context.TODO(), tt.pod, metav1.CreateOptions{})
				if err != nil {
					t.Error(err)
				}
			}
			err = k8sClient.CheckPodsInDefaultNamespace()
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPodsInDefaultNamespace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestKubernetesClient_CheckExposeControlPlane(t *testing.T) {

	clientSet := fake.NewSimpleClientset()
	k8sClient, err := NewKubernetesClient(clientSet)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name      string
		k8sClient KubernetesInterface
		Node      corev1.Node
		wantErr   bool
	}{
		{
			name: "control plane node pass since there is no public IP address for the node",
			Node: corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-private-node",
					Labels: map[string]string{
						"kubernetes.io/hostname": "kind-control-plane",
					},
				},
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    "InternalIP",
							Address: "172.16.10.1",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "control plane node fail if there is a public IP address for the node",
			Node: corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-public-node",
					Labels: map[string]string{
						"kubernetes.io/hostname": "kind-control-plane",
					},
				},
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    "InternalIP",
							Address: "200.190.10.1",
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "non control plane node with public IP without validation",
			Node: corev1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-public-node2",
					Labels: map[string]string{
						"kubernetes.io/hostname": "worker-node",
					},
				},
				Status: corev1.NodeStatus{
					Addresses: []corev1.NodeAddress{
						{
							Type:    "InternalIP",
							Address: "200.190.10.2",
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {

		if len(tt.Node.Status.Addresses) != 0 {
			_, err = clientSet.CoreV1().Nodes().Create(context.Background(), &tt.Node, metav1.CreateOptions{})
			if err != nil {
				t.Error(err)
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			if err := k8sClient.CheckExposeControlPlane(); (err != nil) != tt.wantErr {
				t.Errorf("CheckExposeControlPlane() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		// clean up
		err = clientSet.CoreV1().Nodes().Delete(context.Background(), tt.Node.Name, metav1.DeleteOptions{})
		if err != nil {
			t.Error(err)
		}
	}
}
