#! /bin/sh

PVC_NAME=$1

kubectl patch pvc ${PVC_NAME} -p '{"metadata":{"finalizers": []}}' --type=merge
kubectl delete pvc ${PVC_NAME}