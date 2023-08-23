#!/bin/sh
REPO=$1
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Expose web server
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "NodePort"}}'

wget https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64
install -m 555 argocd-linux-amd64 /usr/local/bin/argocd
rm argocd-linux-amd64

argocd app create sfc --repo ${REPO} --path sfc --dest-server https://kubernetes.default.svc --dest-namespace default