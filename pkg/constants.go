package pkg

const (
	DefaultNamespace    = "it is not recommended to have pods in the default namespace. Please consider moving them to a separate namespace to improve the organization and management of your cluster resources. reference: https://github.com/converge/kind-security-check/blob/main/docs/reference.md#default-namespace-100"
	ExposedControlPlane = "to ensure the security of your Kubernetes control plane, it is recommended to keep it isolated from the internet and not expose it through a public IP address. Consider using a private network or a VPN connection to access the control plane. reference: https://github.com/converge/kind-security-check/blob/main/docs/reference.md#exposed-control-plane-101"
)
