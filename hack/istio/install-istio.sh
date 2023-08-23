#!/bin/sh
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.17.5 TARGET_ARCH=x86_64 sh -
# istioctl operator init
cd istio-${ISTIO_VERSION}
export PATH=$PWD/bin:$PATH
# export KUBECONFIG=
# Init and Install Operator to cluster
./bin/istioctl operator init


# Docs: https://istio.io/latest/docs/setup/install/multicluster/multi-primary_multi-network/
# Set the default network for cluster1
kubectl --context="${CTX_CLUSTER1}" get namespace istio-system && \
kubectl --context="${CTX_CLUSTER1}" label namespace istio-system topology.istio.io/network=network1

# Expose gateway cluster 1 on multiple network
./gen-eastwest-gateway.sh \
    --mesh mesh1 --cluster cluster1 --network network1 | \
    istioctl --context="${CTX_CLUSTER1}" install -y -f -

# Expose services in cluster1
kubectl --context="${CTX_CLUSTER1}" apply -n istio-system -f \
    samples/multicluster/expose-services.yaml

# Set the default network for cluster2
kubectl --context="${CTX_CLUSTER2}" get namespace istio-system && \
kubectl --context="${CTX_CLUSTER2}" label namespace istio-system topology.istio.io/network=network2

# Expose gateway cluster 2 on multiple network
./gen-eastwest-gateway.sh \
    --mesh mesh1 --cluster cluster2 --network network2 | \
    istioctl --context="${CTX_CLUSTER2}" install -y -f -

# Expose services in cluster2
kubectl --context="${CTX_CLUSTER2}" apply -n istio-system -f \
    samples/multicluster/expose-services.yaml

# Enable Endpoint Discovery

# Cluster 1
# Install a remote secret in cluster2 that provides access to cluster1’s API server.
istioctl x create-remote-secret \
  --context="${CTX_CLUSTER1}" \
  --name=cluster1 | \
  kubectl apply -f - --context="${CTX_CLUSTER2}"

# Cluster 2
# Install a remote secret in cluster1 that provides access to cluster2’s API server.
istioctl x create-remote-secret \
  --context="${CTX_CLUSTER2}" \
  --name=cluster2 | \
  kubectl apply -f - --context="${CTX_CLUSTER1}"
# Generate remote secret for cluster 1 connect to cluster 2
 ./istioctl x create-remote-secret --name=cluster1 --kubeconfig ~/.kube/config > access-cluster1.yaml
kubectl apply -f access-cluster2.yaml

# Generate remote secret for cluster 2 connect to cluster 1
 ./istioctl x create-remote-secret --name=cluster2 --kubeconfig ~/.kube/config > access-cluster2.yaml
 kubectl apply -f access-cluster1.yaml
 