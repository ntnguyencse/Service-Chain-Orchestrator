#!/bin/sh

# To begin, create the sample namespace in each cluster:
# kubectl create --context="${CTX_CLUSTER1}" namespace sample
# kubectl create --context="${CTX_CLUSTER2}" namespace sample

# Cluster 1
kubectl create  namespace sample
# Cluster 2
kubectl create namespace sample

# Enable automatic sidecar injection for sample namespace
# kubectl label --context="${CTX_CLUSTER1}" namespace sample \
#     istio-injection=enabled
# kubectl label --context="${CTX_CLUSTER2}" namespace sample \
#     istio-injection=enabled

# Cluster 1
kubectl label namespace sample \
    istio-injection=enabled
# Cluster 2
kubectl label namespace sample \
    istio-injection=enabled

# Create the HelloWorld service in both clusters:
kubectl apply --context="${CTX_CLUSTER1}" \
    -f samples/helloworld/helloworld.yaml \
    -l service=helloworld -n sample
kubectl apply --context="${CTX_CLUSTER2}" \
    -f samples/helloworld/helloworld.yaml \
    -l service=helloworld -n sample

# Create the HelloWorld service in both clusters:
kubectl apply --context="${CTX_CLUSTER1}" \
    -f samples/helloworld/helloworld.yaml \
    -l service=helloworld -n sample
kubectl apply --context="${CTX_CLUSTER2}" \
    -f samples/helloworld/helloworld.yaml \
    -l service=helloworld -n sample
