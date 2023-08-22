#!/bin/sh
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.17.5 TARGET_ARCH=x86_64 sh -
# istioctl operator init
cd istio-${ISTIO_VERSION}
export PATH=$PWD/bin:$PATH
# export KUBECONFIG=
./bin/istioctl operator init



samples/multicluster/gen-eastwest-gateway.sh \
    --mesh mesh1 --cluster cluster1 --network network1 | \
    istioctl --context="${CTX_CLUSTER1}" install -y -f -
