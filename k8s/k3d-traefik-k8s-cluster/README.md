## Create K3D K8s Cluster

To create the K3D K8s cluster issue the following arguments to the `kr8-k3d-traefik-k8s-cluster.sh` script:
The script takes 2 arguments, arg #1 is **# of worker/agent nodes** and arg #2 is **# of server nodes**.

```
sh kr8-k3d-traefik-k8s-cluster.sh 2 1      # Create K3d K8s cluster with 1 server (master) node and 2 agent/worker nodes
```


## Destroy K3D K8s Cluster

To destroy the K3D K8s cluster issue the following:

```
sh kill-k3d-traefik-k8s-cluster.sh
```