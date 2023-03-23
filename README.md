# Rengine Operator

The kubernetes operator of [rengine](https://github.com/wl4g/rengine)

## Getting Started

You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) or [K3S](https://k3s.io/) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Preparing

- 1. Installing kubebuilder

```bash
sudo curl -L -o /bin/kubebuilder 'https://github.com/kubernetes-sigs/kubebuilder/releases/download/v3.9.1/kubebuilder_linux_amd64'
sudo chmod +x /bin/kubebuilder
sudo ln -snf /bin/kubebuilder /bin/kb
```

- 2. Installing certmanager

```bash
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.6.1/cert-manager.yaml
```

### Running on the cluster

- 1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

- 2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=docker.io/wl4g/rengine-operator:latest
```

- 3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=docker.io/wl4g/rengine-operator:latest

kubectl -n rengine-operator-system get po
NAME                                                   READY   STATUS    RESTARTS   AGE
rengine-operator-controller-manager-7567b4fbd9-ngpzp   2/2     Running   0          64s
```

### Uninstall CRDs

To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller

UnDeploy the controller from the cluster:

```sh
make undeploy
```

### Extension module for examples

```bash
# Add api
kb create api --group rengine --version v1beta1 --kind RengineController
# Add webhook
kb create webhook --group rengine --version v1beta1 --kind RengineController \
--defaulting --programmatic-validation
```

### How it works

This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out

- 1. Install the CRDs into the cluster:

```sh
make install
```

- 2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions

If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023 James Wong<jameswong1376@gmail.com>.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

### FAQ

- 1. If you run make docker-build and report an error that gcr.io/distroless/static:nonroot cannot be pulled, you can use the mainland backup image. Note: If you want to make a backup image yourself, you may only be able to pull it in GCP's cloud-shell or VM (because authentication is required on overseas VMs of other cloud vendors)

```bash
docker pull registry.cn-shenzhen.aliyuncs.com/wl4g-k8s/distroless_static:nonroot
docker tag registry.cn-shenzhen.aliyuncs.com/wl4g-k8s/distroless_static:nonroot gcr.io/distroless/static:nonroot
```
