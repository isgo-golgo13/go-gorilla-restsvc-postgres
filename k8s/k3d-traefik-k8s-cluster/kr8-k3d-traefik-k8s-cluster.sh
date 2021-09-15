#! /bin/sh

AGENTS=$1
SERVERS=$2

k3d cluster create k3d-traefik-k8s-cluster \
--agents ${AGENTS} --servers ${SERVERS} \
--api-port 127.0.0.1:6443 \
-p 80:80@loadbalancer \
-p 443:443@loadbalancer \
--k3s-server-arg "--no-deploy=traefik"

helm upgrade --install traefik traefik/traefik