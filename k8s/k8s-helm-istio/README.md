## Table of Contents

The following is table of contents on the topics of how, when and why for defining K8s configs
in the either K8s Helm or K8s Kustomize. The topics will have **1:1** association to a Confluence doc specification providing the required understanding to provision these configs.


| Shell Script        |             Script Description                                                                                           |
| --------------------| -------------------------------------------------------------------------------------------------------------------------|
| **ht.sh**           |  wrapper script issues: **helm template go-gorilla-restsvc <chart-directory>**                                           |
| **hui.sh**          |  wrapper script issues: **helm upgrade --install go-gorilla-restsvc --namespace go-gorilla-restsvc <chart-directory>**   |
| **hu.sh**           |  wrapper script issues: **helm uninstall go-gorilla-restsvc --namespace go-gorilla-restsvc**                             |
| **hu-dryrun.sh**    |  wrapper script issues: **helm uninstall go-gorilla-restsvc --namespace go-gorilla-restsvc <chart-directory> --dry-run** |




## Directory Structure

```
k8s-helm-istio
├── README.md
├── go-gorilla-restsvc
│   ├── Chart.yaml
│   ├── env-preprod-values.yaml
│   ├── env-prod-values.yaml
│   ├── env-staging-values.yaml
│   ├── templates
│   │   ├── _helpers.tpl
│   │   ├── deployment.yaml
│   │   ├── gateway.yaml
│   │   ├── hpa.yaml
│   │   ├── pdb.yaml
│   │   ├── service-account.yaml
│   │   ├── service.yaml
│   │   ├── tests
│   │   │   └── test-connection.yaml
│   │   ├── virtual-service.yaml
│   │   └── vpa.yaml
│   └── values.yaml
├── ht.sh
├── hu-dryrun.sh
├── hu.sh
└── hui.sh
```


## K3D Istio Ingress K8s Cluster Config

Create the K3D K8s cluster with Istio Ingress Controller. The script in `k8s/k3d-istio-k8s-cluster` directory (see `kr8-k3d-istio-k8s-cluster.sh`) will take two args, arg 1: takes the count of agent/worker nodes and arg 2: takes the count of server nodes. 

```
sh kr8-k3d-istio-k8s-cluster.sh 2 1           # create 2 worker/1 server k3d istio k8s cluster
```

This script automatically installs istio into the k3d k8s cluster using profile `demo` (which installs istio ingress-gateway, ingress, egress, istio core and istiod) and label injects the default namespace of the cluster with istio as follows:

```
istioctl install --set profile=demo
kubectl label namespace default istio-injection=enabled
```

## K3D Istio K8s Cluster EnvVars Config (To Allow External App Service Access to K8s Cluster)

```
export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].port}')
export TCP_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="tcp")].port}')

export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT
```






## Pre-Flight Check the Helm Chart (Prior to K8s Helm Chart Install to Cluster)

At the root-level directory where the chart directory is (inside `k8s-helm-istio` directory) which is
the `go-gorilla-restsvc` directory and issue the following helm template wrapper script:

```
sh ht.sh  
```


If the validation is successful the fully rendered (locally rendered) helm templates merged with the values.yaml will parameterize the helm chart to show all the values are plugged in.


### Install the Helm Chart to K8s Cluster 

At the root-level directory where the chart directory is (inside `k8s-helm-istio` directory) which is th
`go-gorilla-restsvc` directory and issue the following helm upgrade/install wrapper script:

```
sh hui.sh
```


** IMPORTANT** : The `hui.sh` helm install/upgrade wrapper script WILL FAIL the first time run as it will indicate the `namespace go-gorilla-restsvc` already exists. This is an oddity in namespace generation in Helm 3 and to deploy successfully just run the script again as there is a delay in auto-generation of the required `go-gorilla-restsvc` namespace. After the first time you run this script and see the error you will see namespace `go-gorilla-restsvc` in the namespace list after doing a `kubectl get namespaces`. Again to successfully deploy this chart into the cluster just run the `sh hui.sh` script a second time. You should see the chart deployment confirmation. This `hui.sh` helm wrapper script includes the `--create-namespace` flag at the end of the `helm install --upgrade <HELM-RELEASE-NAME> <HELM-CHART-DIRECTORY>` command.



## Access the Service 

Access the `/health-check` service endpoint (handler is HealthCheck in the logs)
```
curl -s "http://${GATEWAY_URL}/health-check"       # Returns Status 200 and "Service Healthy"
```

The result of this service SHOULD show:
```
"Service Healthy"
```

Access the `/engines` service endpoint (handler is Engines in the logs)
```
curl -s "http://${GATEWAY_URL}/engines"           # Returns array of automotive engines
```

The result of this service endpoint SHOULD show:
```
[{"id":"100000001","serial_id":"VW_100000001","engine_config":"V8","engine_capacity":250.5},{"id":"100000002","serial_id":"Audi_100000002","engine_config":"V8","engine_capacity":220.5},{"id":"100000003","serial_id":"Porsche_100000003","engine_config":"V8","engine_capacity":50.5},{"id":"100000004","serial_id":"Porsche_100000004","engine_config":"V8","engine_capacity":270.5},{"id":"100000005","serial_id":"Mercedes_10000005","engine_config":"V8-Twin-Turbo","engine_capacity":250.25},{"id":"100000006","serial_id":"Mercedes_10000006","engine_config":"V8-Twin-Turbo","engine_capacity":270.25},{"id":"100000007","serial_id":"Mercedes_10000007","engine_config":"V12-Twin-Turbo","engine_capacity":350.75},{"id":"100000008","serial_id":"Mercedes_10000008","engine_config":"V8","engine_capacity":250.5}]
```

Access the `engines/{id}` service endpoint (handler is Engine in the logs)
```
curl -s "http://${GATEWAY_URL}/engines/4"         # Returns nth automotive engine in the storage array (bounds checked)
```

The result of this service endpoint SHOULD show:
```
{"id":"100000005","serial_id":"Mercedes_10000005","engine_config":"V8-Twin-Turbo","engine_capacity":250.25}
```




### Uninstall the Helm Chart from K8s Cluster 

At the root-level directory where the chart directory is (inside `k8s-helm-istio` directory) which is th
`go-gorilla-restsvc` directory and issue the following helm uninstall wrapper script:

```
sh hu.sh
```