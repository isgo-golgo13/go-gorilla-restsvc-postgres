#! /bin/sh

#First time this script is run will result in error "namespaces go-gorilla-restsvc already exists" to fix run this script again and will get deployed
helm upgrade --install go-gorilla-restsvc-postgres --namespace go-gorilla-restsvc-postgres ./go-gorilla-restsvc-postgres --create-namespace