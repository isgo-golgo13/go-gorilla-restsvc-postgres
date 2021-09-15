#! /bin/sh

AGENTS=$1
SERVERS=$2

k3d cluster create k3d-istio-k8s-cluster \
--agents ${AGENTS} --servers ${SERVERS} \
--api-port 127.0.0.1:6443 \
-p 80:80@loadbalancer \
-p 443:443@loadbalancer \
--k3s-server-arg "--no-deploy=traefik"

istioctl install --set profile=demo

kubectl label namespace default istio-injection=enabled