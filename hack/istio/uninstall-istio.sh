#!/bin/sh

istioctl uninstall  -y --purge
kubectl delete ns istio-system 