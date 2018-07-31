# Scripts to Install Docker, Kubernetes, Helm and Istio on Ubuntu

## Create k8s master node via kubeadmin
`1_install_k8s_master.sh`

## Create k8s work node via kubeadmin
`2_install_k8s_minion.sh`

## Install Istio via helm
`3_install_istio.sh`

## Lable the namespaces in which you want to enable auto sidecar injection
`kubectl label namespace default istio-injection=enabled`
  
## Notice
Sidecar auto injection is disabled, so the sidecar injector will not inject the sidecar into pods by default. Add the sidecar.istio.io/inject annotation with value true to the pod template spec to enable injection.

The following example uses the sidecar.istio.io/inject annotation to enable sidecar injection.
```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: ignored
spec:
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "true"
    spec:
      containers:
      - name: ignored
        image: tutum/curl
        command: ["/bin/sleep","infinity"]
```

You can enable sidecar auto injection by setting the injection policy to 'enabled' at line 835 of istio.yaml.
```
 822 apiVersion: v1
 823 kind: ConfigMap
 824 metadata:
 825   name: istio-sidecar-injector
 826   namespace: istio-system
 827   labels:
 828     app: istio
 829     chart: istio-0.8.0
 830     release: RELEASE-NAME
 831     heritage: Tiller
 832     istio: sidecar-injector
 833 data:
 834   config: |-
 835     policy: disabled
 836     template: |-
 837       initContainers:
 838       - name: istio-init
 839         image: docker.io/istio/proxy_init:0.8.0
```

For more information on Istio integration, refer to:

 -  [Manage ONAP Microservices with Istio Service Mesh](https://wiki.onap.org/display/DW/Manage+ONAP+Microservices+with+Istio+Service+Mesh)
