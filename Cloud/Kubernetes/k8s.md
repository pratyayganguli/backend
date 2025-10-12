### Kubernetes Basic to Advanced

#### Kubernetes: #### 
Kubernetes is a container orchestration developed by Google, it automates deployment, scaling, load balancing of containers. It acts like an os for containerized applications.

#### Architecture of Kubernetes:

- **Control Plane-** The control plane acts like a brain of the kubernetes cluster.
    - **Kube API Server:** This is the entry point of the Control Plane where, we can communicate with the kubernetes REST apis. Whenever we use the `kubectl` command it goes through the Kube API Server.
    - **etcd-** A distrubted key-value pair data storage which stores and manages the data related to the cluster.
    - **kube-scheduler:** It determines in which node a pod should run, based on the resource requiremnets and other determining factors.
    - **kube-controller-manager:** It is responsible for running controller loops for determining the cluster's current state vs desired state.
     which
- **Worker Node Components-**
    - **Kubelet:** An agent that works on each worker node, communicates with the API server
    - **Kube-proxy:** It is responsible for handling networking for pods on the node and mantains network rules to allow communication between pods and services.
    - **Container Runtime:** This component is resposible for running the containers in a node.

#### Helm -  #### 
Helm acts like a package manager for Kubernetes, manages all the kubernetes manifests.
