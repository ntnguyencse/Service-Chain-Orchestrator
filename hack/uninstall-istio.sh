#!/bin/sh

istioctl uninstall --context="${CTX_CLUSTER}" -y --purge
kubectl delete ns istio-system --context="${CTX_CLUSTER}"
