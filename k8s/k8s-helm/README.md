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
tree k8s-helm                                                                  (master|…1)
k8s-helm
├── README.md
├── go-gorilla-restsvc
│   ├── Chart.yaml
│   ├── README.md
│   ├── env-preprod-values.yaml
│   ├── env-prod-values.yaml
│   ├── env-staging-values.yaml
│   ├── templates
│   │   ├── _helpers.tpl
│   │   ├── deployment.yaml
│   │   ├── hpa.yaml
│   │   ├── ingress.yaml
│   │   ├── pdb.yaml
│   │   ├── service.yaml
│   │   ├── serviceaccount.yaml
│   │   ├── tests
│   │   │   └── test-connection.yaml
│   │   └── vpa.yaml
│   └── values.yaml
├── ht.sh
├── hu-dryrun.sh
├── hu.sh
└── hui.sh

```




#  K8s Deployment using K8s Helm 3

### Pre-Flight Check the Helm Chart (Prior to K8s Helm Chart Install to Cluster)

At the root-level directory where the chart directory is (inside `k8s-helm` directory) which is
the `go-gorilla-restsvc` directory and issue the following helm template wrapper script:

```
sh ht.sh  
```


If the validation is successful the fully rendered (locally rendered) helm templates merged with the values.yaml will parameterize the helm chart to show all the values are plugged in.


### Install the Helm Chart to K8s Cluster 

At the root-level directory where the chart directory is (inside `k8s-helm` directory) which is th
`go-gorilla-restsvc` directory and issue the following helm upgrade/install wrapper script:

```
sh hui.sh
```


### Execute the service (service active in K8s running through the Ingress)

#### Get Health Check
```
curl localhost/health-check
```

#### Get all Engines
```
curl localhost/engines
```

#### Get an Engine [0th Engine - Engines -1]
```
curl localhost/engine/2
```



### Uninstall the Helm Chart from K8s Cluster 

At the root-level directory where the chart directory is (inside `k8s-helm` directory) which is th
`go-gorilla-restsvc` directory and issue the following helm uninstall wrapper script:

```
sh hu.sh
```