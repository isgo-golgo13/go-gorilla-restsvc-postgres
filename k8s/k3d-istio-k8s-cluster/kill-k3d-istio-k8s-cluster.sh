#! /bin/sh

istioctl x uninstall --purge
k3d cluster delete k3d-istio-k8s-cluster
