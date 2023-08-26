

# Install istio 

istioctl install


helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update


helm install edge-cluster-3 prometheus-community/kube-prometheus-stack

kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Expose web server
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "NodePort"}}'
