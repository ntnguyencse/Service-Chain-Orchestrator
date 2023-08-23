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
# kubectl apply --context="${CTX_CLUSTER1}" \
#     -f samples/helloworld/helloworld.yaml \
#     -l service=helloworld -n sample
# kubectl apply --context="${CTX_CLUSTER2}" \
#     -f samples/helloworld/helloworld.yaml \
#     -l service=helloworld -n sample

# Create the HelloWorld service in both clusters:
kubectl apply -f helloworld.yaml \
    -l service=helloworld -n sample
kubectl apply -f samples/helloworld/helloworld.yaml \
    -l service=helloworld -n sample

# Deploy v1 on cluster 1
kubectl apply  \
    -f helloworld.yaml \
    -l version=v1 -n sample

# Deploy v2 on cluster 2
kubectl apply  \
    -f helloworld.yaml \
    -l version=v2 -n sample


# Delpoy sleep on both clusters
kubectl apply \
    -f sleep.yaml -n sample
kubectl apply  \
    -f sleep.yaml -n sample


# Verify 
kubectl exec  -n sample -c sleep \
    "$(kubectl get pod  -n sample -l \
    app=sleep -o jsonpath='{.items[0].metadata.name}')" \
    -- curl -sS helloworld.sample:5000/hello

