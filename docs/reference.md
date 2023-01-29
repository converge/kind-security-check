
## Default namespace (100)

It is generally recommended not to deploy applications into the default Kubernetes namespace for a number of reasons:

1. Namespaces provide a way to divide cluster resources among multiple users or teams. Using the default namespace can lead to resource contention and make it difficult to manage and monitor different resources.
2. The default namespace is used for system-level resources, such as the Kubernetes control plane components. Mixing system-level resources with application-level resources can create confusion and make it more difficult to troubleshoot issues.
3. Having multiple namespaces allows for better isolation of resources and can help with security. Resources in different namespaces cannot interact with each other unless explicitly granted permissions.
4. By deploying applications into separate namespaces, it's possible to use different resource quotas and limit ranges, which can help ensure that a single application or namespace doesn't consume all the resources in the cluster.
5. It would be easier to manage and delete resources if they are in different namespaces. Instead of deleting all resources in the cluster, you can just delete the resources in a specific namespace.
It's best practice to create a new namespace for each application or team and deploy the resources there.

## Exposed control plane (101)

Kubernetes control plane, which includes components such as the API server, etcd, and controller manager, should not have a public IP address for security reasons.

The control plane is responsible for managing the state of the cluster, and it exposes various APIs that can be used to interact with the cluster. If the control plane is accessible from the internet, it increases the attack surface of the cluster and makes it vulnerable to various types of attacks.

Additionally, exposing the control plane to the internet can also create a risk of data breaches, as sensitive information such as secrets and configurations are stored in the control plane.

It is generally recommended to deploy the control plane behind a firewall and to use secure communication channels, such as VPN or SSH tunnels, to access it.

Also, It's a best practice to have a control plane as a private IP address, behind a firewall, and only accessible to authorized cluster administrators.
