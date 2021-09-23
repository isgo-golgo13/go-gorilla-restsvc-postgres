#! /bin/sh

helm uninstall go-gorilla-restsvc-postgres --namespace go-gorilla-restsvc-postgres
kubectl delete namespace go-gorilla-restsvc-postgres 