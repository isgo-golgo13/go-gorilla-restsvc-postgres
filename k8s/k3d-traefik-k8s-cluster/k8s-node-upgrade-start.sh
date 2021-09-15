#! /bin/sh

K8S_TARGET_DRAIN_NODE=$1

kubectl drain ${K8S_TARGET_DRAIN_NODE} --force  
