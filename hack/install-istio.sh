#!/bin/sh
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.17.5 TARGET_ARCH=x86_64 sh -
istioctl operator init
cd istio-${ISTIO_VERSION}
export PATH=$PWD/bin:$PATH
# export KUBECONFIG=
./bin/istioctl operator init
