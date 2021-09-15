#! /bin/sh

K8S_TARGET_UNCORDON_NODE=$1

kubectl uncordon ${K8S_TARGET_UNCORDON_NODE}  