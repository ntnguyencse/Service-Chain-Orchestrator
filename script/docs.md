# Installing and Seting up L-KaaS Project

## Prequisite
A Kubernetes for management and environment to running workload.
Must installed these software:
- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### Install kubebuilder:


#### Download kubebuilder and install locally.

    curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
    chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
## Run project

### Regenerate CRDs
If you are editing the API definitions, generate the manifests such as Custom Resources (CRs) or Custom Resource Definitions (CRDs) using


    make manifests
### Install Instances of Custom Resources

If you pressed y for Create Resource [y/n] then you created a CR for your CRD in your samples (make sure to edit them first if you’ve changed the API definition):

    kubectl apply -k config/samples/

### Run your controller 
Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):


    make run
### Connect to K8s Cluster
You’ll need a Kubernetes cluster to run against. You can use KIND to get a local cluster for testing, or run against a remote cluster.

*Context Used*
Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster kubectl cluster-info shows).

Install the CRDs into the cluster:


    make install

### Install CAPI

Prequisites: 
Common Prerequisites
Install and setup kubectl in your local environment
Install kind and Docker
Install Helm

#### Initialize cluster
Initialize the management cluster
Now that we’ve got clusterctl installed and all the prerequisites in place, let’s transform the Kubernetes cluster into a management cluster by using clusterctl init.

The command accepts as input a list of providers to install; when executed for the first time, clusterctl init automatically adds to the list the cluster-api core provider, and if unspecified, it also adds the kubeadm bootstrap and kubeadm control-plane providers.

    curl -L <https://github.com/kubernetes-sigs/cluster-api/releases/download/v1.4.3/clusterctl-linux-amd64> -o clusterctl
    sudo install -o root -g root -m 0755 clusterctl /usr/local/bin/clusterctl
    clusterctl version
    export CLUSTER_TOPOLOGY=true

    `#Initialize the management cluster`

    clusterctl init --infrastructure=openstack:v0.6.4 --core=cluster-api:v1.4.0 --bootstrap=kubeadm:v1.4.0 --control-plane=kubeadm:v1.4.0 -v5

    clusterctl init --infrastructure=openstack:v0.7.0,aws:v2.2.2 --core=cluster-api:v1.4.0 --bootstrap=kubeadm:v1.4.0 --control-plane=kubeadm:v1.4.0 -v5

    `# clusterctl init --infrastructure=openstack:v0.7.0 -v5`

    clusterctl init --infrastructure=openstack:v0.7.0,aws:v2.2.2 --core=cluster-api:v1.4.0 --bootstrap=kubeadm:v1.4.0 --control-plane=kubeadm:v1.4.0 --config /home/ubuntu/aws-capi/capi-config/config.yaml -v5


## Configure system

### Cluster API Installation

Install clusterctl as follow:

```
    curl -L https://github.com/kubernetes-sigs/cluster-api/releases/download/v1.4.3/clusterctl-linux-amd64 -o clusterctl
    sudo install -o root -g root -m 0755 clusterctl /usr/local/bin/clusterctl
    clusterctl version
    export CLUSTER_TOPOLOGY=true

```

**Avoiding GitHub rate limiting**
While using providers hosted on GitHub, clusterctl is calling GitHub API which are rate limited; for normal usage free tier is enough but when using clusterctl extensively users might hit the rate limit.

To avoid rate limiting for the public repos set the GITHUB_TOKEN environment variable. To generate a token follow this documentation. The token only needs repo scope for clusterctl.

Per default clusterctl will use a go proxy to detect the available versions to prevent additional API calls to the GitHub API. It is possible to configure the go proxy url using the GOPROXY variable as for go itself (defaults to https://proxy.golang.org). To immediately fallback to the GitHub client and not use a go proxy, the environment variable could get set to GOPROXY=off or GOPROXY=direct. If a provider does not follow Go’s semantic versioning, clusterctl may fail when detecting the correct version. In such cases, disabling the go proxy functionality via GOPROXY=off should be considered.

### Build clusterctl config

Create template config file:
```
    #!/usr/bin/env bash
    cat << EOF > clusterctl-config.yml
    OPENSTACK_CLOUD: $OPENSTACK_CLOUD
    OPENSTACK_CLOUD_CACERT_B64: $OPENSTACK_CLOUD_CACERT_B64
    OPENSTACK_CLOUD_PROVIDER_CONF_B64: $OPENSTACK_CLOUD_PROVIDER_CONF_B64
    OPENSTACK_CLOUD_YAML_B64: $OPENSTACK_CLOUD_YAML_B64

    # The list of nameservers for OpenStack Subnet being created

    # Set this value when you need create a new network/subnet while the access through DNS is required

    OPENSTACK_DNS_NAMESERVERS: 8.8.8.8

    # FailureDomain is the failure domain the machine will be created in

    OPENSTACK_FAILURE_DOMAIN: compute

    # OPENSTACK_FAILURE_DOMAIN: $(OPENSTACK_FAILURE_DOMAIN)

    # The flavor reference for the flavor for your server instance

    OPENSTACK_CONTROL_PLANE_MACHINE_FLAVOR: m1.medium

    # The flavor reference for the flavor for your server instance

    OPENSTACK_NODE_MACHINE_FLAVOR: m1.medium

    # The name of the image to use for your server instance. If the RootVolume is specified, this will be ignored and use rootVolume directly

    OPENSTACK_IMAGE_NAME: ubuntu-k8-1.22

    # The SSH key pair name

    OPENSTACK_SSH_KEY_NAME: $OPENSTACK_SSH_KEY_NAME

    # The external network

    OPENSTACK_EXTERNAL_NETWORK_ID: $OPENSTACK_EXTERNAL_NETWORK_ID

    # Enabling Feature Gates

    CLUSTER_TOPOLOGY: true

    # Kubernetes Version

    KUBERNETES_VERSION: $KUBERNETES_VERSION

    # CLUSTER_TOPOLOGY=true

    EOF
```

Output similar like this:
    The output of clusterctl init is similar to this:
```
    Fetching providers
    Installing cert-manager Version="v1.11.0"
    Waiting for cert-manager to be available...
    Installing Provider="cluster-api" Version="v1.0.0" TargetNamespace="capi-system"
    Installing Provider="bootstrap-kubeadm" Version="v1.0.0" TargetNamespace="capi-kubeadm-bootstrap-system"
    Installing Provider="control-plane-kubeadm" Version="v1.0.0" TargetNamespace="capi-kubeadm-control-plane-system"
    Installing Provider="infrastructure-docker" Version="v1.0.0" TargetNamespace="capd-system"

    Your management cluster has been initialized successfully!

    You can now create your first workload cluster by running the following:

    clusterctl generate cluster [name] --kubernetes-version [version] | kubectl apply -f -

```


#### Install ClusterAPI Provider

The clusterctl init command installs the Cluster API components and transforms the Kubernetes cluster into a management cluster.

This document provides more detail on how `clusterctl init` works and on the supported options for customizing your management cluster.
```
    # Initialize the management cluster
    clusterctl init --infrastructure=openstack:v0.7.0,aws:v2.2.2 --core=cluster-api:v1.4.0 --bootstrap=kubeadm:v1.4.0 --control-plane=kubeadm:v1.4.0 --config /home/ubuntu/aws-capi/capi-config/config.yaml -v5
```
*You can specify the provider version by appending a version tag to the provider name, e.g. aws:v0.4.1.*

#### Prepare Kubernetes cluster image for Cluster API on OpenStack

If we want to use Openstack provider, we need to build from source or use pre-build image for Kubernetes cluster.
The tool is used for building iamge is Image Builder: The Image Builder *(https://image-builder.sigs.k8s.io/capi/capi.html)* can be used to build images intended for use with Kubernetes CAPI providers. Each provider has its own format of images that it can work with. For example, AWS instances use AMIs, and vSphere uses OVAs.
##### Build from source

A ClusterAPI compatible image must be available in your OpenStack. For instructions on how to build a compatible image see image-builder. Depending on your OpenStack and underlying hypervisor the following options might be of interest:

image-builder (OpenStack) 
`https://image-builder.sigs.k8s.io/capi/providers/openstack.html`
image-builder (vSphere)
`https://image-builder.sigs.k8s.io/capi/providers/vsphere.html`

##### Using pre-build Kubernetes image

We collected several pre-build images and stored in `\\114.71.50.192\personal\47.NguyenThanhNguyen\openstack-image`

Original sources: `https://github.com/osism/k8s-capi-images`
There are tested images and could be directly use woth Openstack with few steps. Upload these image to Openstack before creating Kubernetes cluster.

## Install Helm

Install Kubernetes or have access to a cluster
You must have Kubernetes installed. For the latest release of Helm, we recommend the latest stable release of Kubernetes, which in most cases is the second-latest minor release.
You should also have a local configured copy of kubectl.

```
    # Download a installation script of the Helm client.
    curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
    chmod 700 get_helm.sh
    ./get_helm.sh
    # Or directly download a binary release of the Helm client. 
    curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```


## Configure L-KaaS system

There are 3 type of configurations in L-KaaS:
```
const DEFAULT_CONFIG_PATH string = "./config.yml"
const DEFAULT_OPENSTACKCONFIG_PATH string = "./openstack-config.yml"
const DEFAULT_AWSCONFIG_PATH string = "./capactl.yml"
const DEFAULT_CAPI_CONFIG_PATH string = "/.l-kaas/config/capi/capictl.yml"
const DEFAULT_CAPI_AWS_CONFIG_PATH string = "/.l-kaas/config/capi/capactl.yml"
```

### Get OpenStack credentials

To set the required environment variables for the OpenStack command-line clients, you must create an environment file called an OpenStack rc file, or openrc.sh file. If your OpenStack installation provides it, you can download the file from the OpenStack Dashboard as an administrative user or any other user. This project-specific environment file contains the credentials that all OpenStack services use.

**Download and source the OpenStack RC file**

1. Log in to the dashboard and from the drop-down list select the project for which you want to download the OpenStack RC file.

2. On the Project tab, open the Compute tab and click Access & Security.

3. On the API Access tab, click Download OpenStack RC File and save the file. The filename will be of the form PROJECT-openrc.sh where PROJECT is the name of the project for which you downloaded the file.

4. Copy the PROJECT-openrc.sh file to the computer from which you want to run OpenStack commands.

*For example:* copy the file to the computer from which you want to upload an image with a glance client command.

On any shell from which you want to run OpenStack commands, source the PROJECT-openrc.sh file for the respective project.

### Configure clusterctl configuration file

To see all required OpenStack environment variables execute:

```
clusterctl generate cluster --infrastructure openstack --list-variables capi-quickstart
```

The following script can be used to export some of them:

```
wget https://raw.githubusercontent.com/kubernetes-sigs/cluster-api-provider-openstack/master/templates/env.rc -O /tmp/env.rc
source /tmp/env.rc <path/to/clouds.yaml> <cloud>
```
Apart from the script, the following OpenStack environment variables are required.

```
# The list of nameservers for OpenStack Subnet being created.
# Set this value when you need create a new network/subnet while the access through DNS is required.
export OPENSTACK_DNS_NAMESERVERS=<dns nameserver>
# FailureDomain is the failure domain the machine will be created in.
export OPENSTACK_FAILURE_DOMAIN=<availability zone name>
# The flavor reference for the flavor for your server instance.
export OPENSTACK_CONTROL_PLANE_MACHINE_FLAVOR=<flavor>
# The flavor reference for the flavor for your server instance.
export OPENSTACK_NODE_MACHINE_FLAVOR=<flavor>
# The name of the image to use for your server instance. If the RootVolume is specified, this will be ignored and use rootVolume directly.
export OPENSTACK_IMAGE_NAME=<image name>
# The SSH key pair name
export OPENSTACK_SSH_KEY_NAME=<ssh key pair name>
# The external network
export OPENSTACK_EXTERNAL_NETWORK_ID=<external network ID>

```

A full configuration reference can be found in `https://github.com/kubernetes-sigs/cluster-api-provider-openstack/blob/main/docs/book/src/clusteropenstack/configuration.md`.

### Generate a cluster for test

```
clusterctl generate cluster test \
  --infrastructure openstack \
  --kubernetes-version v1.28.0 \
  --control-plane-machine-count=3 \
  --worker-machine-count=3 \
  --configure /home/ubuntu/l-kaas/L-KaaS/config/capi/clusterctl-config.yaml \
  > capi-quickstart.yaml


clusterctl generate cluster my-cluster --kubernetes-version v1.28.0 \
    --infrastructure aws --config /home/ubuntu/aws-capi/capi-config/config.yaml > my-cluster.yaml

### Openstack
clusterctl generate cluster my-cluster --kubernetes-version v1.28.0     --infrastructure openstack:v0.7.0  --config /home/ubuntu/aws-capi/capi-config/config.yaml > my-cluster-openstack.yaml


### AWS
clusterctl generate cluster my-cluster --kubernetes-version v1.28.0     --infrastructure aws:v2.2.2  --config /home/ubuntu/aws-capi/capi-config/config.yaml > my-cluster-aws.yaml
```


### Making Infrastructure Profiles

Infrastructure Profile Resource referred to pre-defined configurations that contains configurations about Cloud provider, metadata, cluster settings,…
To create `Infrastructure Profiles` we need to research and investigate the configuration of corresponding provider. 
Example: with Openstack provider
There is a github repository contains source code of OpenStack provider in `https://github.com/kubernetes-sigs/cluster-api-provider-openstack`
In `https://github.com/kubernetes-sigs/cluster-api-provider-openstack/tree/main/templates` there a template and configuration for provisioning Kubernetes cluster of Openstack provider. Follow template, we could collect customizable configurations for customzing cluster and put it in to Infrastructure Profiles.

Example: Cluster API and Openstack provider profiles
```
# Cluster API Profile
apiVersion: intent.automation.dcn.ssu.ac.kr/v1
kind: Profile
metadata:
  labels:
    kind: Infrastructure
    namespace: default
    revision: 1.0.0
    type: ClusterAPITemplate
  name: clusterapi-resource-template-default
  namespace: default
spec:
  values:
    filename: cluster-template.yaml
    url: https://github.com/ntnguyencse/L-KaaS/blob/dev/templates/capi/
    version: v0.7.0
---
# Openstack Profiles
apiVersion: intent.automation.dcn.ssu.ac.kr/v1
kind: Profile
metadata:
  labels:
    kind: Infrastructure
    name: openstack-small-cluster
    provider: openstack
    revision: 1.0.0
    type: ClusterTemplate
  name: openstack-small-cluster
  namespace: default
spec:
  blueprint:
  - name: clusterapi-resource-template-default
    revision: 1.0.0
    type: ClusterAPITemplate
  values:
    controlPlaneMachineCount: "1"
    controlplaneFlavor: cluster.controller
    kubernetesVersion: v1.24.8
    workerFlavor: cluster.controller
    workerMachineCount: "1"
```

### Convert CAPI Configurations to L-KaaS Configurations

Because L-KaaS need a configurations for clusterctl client inside of controllers, we need to convert CAPI configuration to format of L-KaaS as follow:

```
#!/usr/bin/env bash
cat << EOF > capictl-config.yml
openstackimagename: $OPENSTACK_IMAGE_NAME
openstackexternalnetworkid: $OPENSTACK_IMAGE_NAME
openstackdnsnameservers: $OPENSTACK_IMAGE_NAME
openstacksshkeyname: $OPENSTACK_IMAGE_NAME
openstackcloudcacertb64: $OPENSTACK_IMAGE_NAME
openstackcloudproviderconfb64: $OPENSTACK_IMAGE_NAME
openstackcloudyamlb64: $OPENSTACK_IMAGE_NAME
openstackfailuredomain: $OPENSTACK_IMAGE_NAME
openstackcloud:  $OPENSTACK_IMAGE_NAME
openstackcontrolplanemachineflavor: $OPENSTACK_IMAGE_NAME
openstacknodemachineflavor : $OPENSTACK_IMAGE_NAME
kubernetesversion: $OPENSTACK_IMAGE_NAME
EOF
```

*Value of configuration is corresponding to CAPI configurations*

*Example:*
capictl-config.yml
```yaml
openstackimagename: "ubuntu-k8s-1.24"
openstackexternalnetworkid: "<network-id>"
openstackdnsnameservers: "8.8.8.8"
openstacksshkeyname: <name>
openstackcloudcacertb64: "<cacert64>"
openstackcloudproviderconfb64: "<value>"
openstackcloudyamlb64: "<value>"
openstackfailuredomain: "compute"
openstackcloud:  "openstack"
openstackcontrolplanemachineflavor: "m1.medium"
openstacknodemachineflavor : "cluster.compute.small"
kubernetesversion: "1.24.8"
```
CAPI Configs: 
```bash
OPENSTACK_CLOUD: openstack
OPENSTACK_CLOUD_CACERT_B64: <cacert-b64>
OPENSTACK_CLOUD_PROVIDER_CONF_B64: <conf-b64>
OPENSTACK_CLOUD_YAML_B64: <yaml-b64>
# The list of nameservers for OpenStack Subnet being created.
# Set this value when you need create a new network/subnet while the access through DNS is required.
OPENSTACK_DNS_NAMESERVERS: 8.8.8.8
# FailureDomain is the failure domain the machine will be created in.
OPENSTACK_FAILURE_DOMAIN: compute
# OPENSTACK_FAILURE_DOMAIN: compute2/compute3
# The flavor reference for the flavor for your server instance.
OPENSTACK_CONTROL_PLANE_MACHINE_FLAVOR: m1.medium
# The flavor reference for the flavor for your server instance.
OPENSTACK_NODE_MACHINE_FLAVOR: cluster.compute.small
# The name of the image to use for your server instance. If the RootVolume is specified, this will be ignored and use rootVolume directly.
OPENSTACK_IMAGE_NAME: ubuntu-k8s-1.24
# The SSH key pair name
OPENSTACK_SSH_KEY_NAME: <ssh-name>
# The external network
OPENSTACK_EXTERNAL_NETWORK_ID: <network-id>
# Enabling Feature Gates
CLUSTER_TOPOLOGY: true
# Kubernetes Version
KUBERNETES_VERSION: 1.24.8
# CLUSTER_TOPOLOGY=true
```
Openstack cloud file:

```bash
# This is a clouds.yaml file, which can be used by OpenStack tools as a source
# of configuration on how to connect to a cloud. If this is your only cloud,
# just put this file in ~/.config/openstack/clouds.yaml and tools like
# python-openstackclient will just work with no further config. (You will need
# to add your password to the auth section)
# If you have more than one cloud account, add the cloud entry to the clouds
# section of your existing file and you can refer to them by name with
# OS_CLOUD=openstack or --os-cloud=openstack
clouds:
  openstack:
    auth:
      auth_url: http://<Openstack-server>/identity
      username: <username>
      password: <password>
      project_id: <project-id>
      project_name: <project-name>
      user_domain_name: <domain-name>
    region_name: <region-name>
    identity_api_version: <version>
    volume_api_version: <version>

```

### Get AWS Credentials

Path of CAPA (Cluster API for AWS): `config/capi/capactl.yaml`
Example configurations for clusterctl client inside L-KaaS: 

```bash
    awsregion: "ap-northeast-2"
    awssshkeyname: "Nguyen"
    awscontrolplanemachinetype: "t3.large"
    awsnodemachinetype: "t3.large"
    kubernetesversion: "1.24.8"
    awsb64encodedcredentials: "aaaaa"
```



### Frequent Bugs

#### Checking Logs of CAPI provider

The Cluster API project is committed to improving the SRE/developer experience when troubleshooting issues, and logging plays an important part in this goal.

**In Cluster API we strive to follow three principles while implementing logging:**

- Logs are for SRE & developers, not for end users! Whenever an end user is required to read logs to understand what is happening in the system, most probably there is an opportunity for improvement of other observability in our API, like e.g. conditions and events.
- Navigating logs should be easy: We should make sure that SREs/Developers can easily drill down logs while investigating issues, e.g. by allowing to search all the log entries for a specific Machine object, eventually across different controllers/reconciler logs.
- Cluster API developers MUST use logs! As Cluster API contributors you are not only the ones that implement logs, but also the first users of them. Use it! Provide feedback!

```bash
#!/usr/bin/env bash

export KUBECONFIG=~/config
# See log of openstack controller
kubectl logs -f -l=cluster.x-k8s.io/provider=infrastructure-openstack -n capo-system
# See log of cluster api controller
kubectl logs -f -l=cluster.x-k8s.io/provider=infrastructure-openstack -n capi-system
```
# Logical Kubernetes as a Service Docs

## User Guide and Design

`https://github.com/ntnguyencse/L-KaaS/tree/main/docs`

## Installing Software

## EMCO

### EMCO Installation

Using Helm to install: `https://gitlab.com/project-emco/core/emco-base/-/blob/main/docs/user/install/Tutorial_Helm.md?ref_type=heads`

For local Installation:
Follow this url: `https://gitlab.com/project-emco/core/emco-base/-/blob/main/docs/user/install/Tutorial_Local_Install.md?ref_type=heads`


### EMCO Configurations

Default EMCO Configuration fileoath: 
`EMCO_DEFAULT_CONFIG_FILE_PATH          = "/.l-kaas/config/emco/.emco.yaml"`
Below is a template that is using by emcoctl

Example configuration of local emcoctl: 
```yaml
  orchestrator:
    host: 192.168.40.15
    port: 30415
  clm:
    host: 192.168.40.15
    port: 30461
  ncm:
    host: 192.168.40.15
    port: 30481
  ovnaction:
    host: 192.168.40.15
    port: 30451
  dcm:
    host: 192.168.40.15
    port: 30477
  gac:
    host: 192.168.40.15
    port: 30420
  dtc:
    host: 192.168.40.15
    port: 30418

```

### Git 

Git configuration file:

Before running controllers we need to export Githu credentials:
```bash
    #!/bin/bash
    export KUBECONFIG=/home/ubuntu/aws-capi/test-capi
    export GH_TOKEN=< Github ID token (write and read permission)>
    export GITHUB_TOKEN=<Github token>
```

### Argo CD

Using script:

```
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
```

### Prometheus

Using helm client to install Prometheus to Kubernetes cluster.

```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update


helm install edge-cluster-1 prometheus-community/kube-prometheus-stackhelm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update


helm install edge-cluster-1 prometheus-community/kube-prometheus-stack
```
# FAQ

#### Stucked Objects

 You may experience, Kubernetes object deletion get stuck, we need to perform some actions to delete forcefully. In this will see how we can remove the finalizer and clean the Kubernetes objects.
```
#!/usr/bin/env bash
NAMESPACE = "namespace"
CLUSTER_NAME = "cluster_name"
kubectl patch crd/clusters.cluster.x-k8s.io -p ${CLUSTER_NAME} '{"metadata":{"finalizers":[]}}' --type=merge -n $NAMESPACE
kubectl patch crd/clusterclasses.cluster.x-k8s.io -p   '{"metadata":{"finalizers":[]}}' --type=merge -n $NAMESPACE
kubectl patch crd/machinedeployments.cluster.x-k8s.io -p  '{"metadata":{"finalizers":[]}}' --type=merge -n $NAMESPACE
kubectl patch crd/openstackmachines.infrastructure.cluster.x-k8s.io  -p  '{"metadata":{"finalizers":[]}}' --type=merge -n $NAMESPACE 
kubectl patch crd/openstackmachinetemplates.infrastructure.cluster.x-k8s.io -p ${CLUSTER_NAME}-md-0 '{"metadata":{"finalizers":[]}}' --type=merge -n $NAMESPACE
```

**Why object still in deleting state**

Kubernetes has its own way of managing memory and resources so does its own Garbage Collection System. It is a systematic way to remove unused/unutilized space. Programming Languages like Java/GO and the Servers built on them all have this process to ensure the memory is managed optimally.

Now, Kubernetes being the modern solution, It does manage its resources and perform Garbage collections when the resources are deleted and in various other contexts too.

Now lets come back to our Kubectl delete and why we are talking about Garbage Collection now.

Kubernetes adds a special tag or annotation to the resource called Finalizers when it creates resources which have multiple dependencies. Here is the definition able Finalizers.

> FINALIZERS ARE NAMESPACED KEYS THAT TELL KUBERNETES TO WAIT UNTIL SPECIFIC CONDITIONS ARE MET BEFORE IT FULLY DELETES RESOURCES MARKED FOR DELETION. FINALIZERS ALERT CONTROLLERS TO CLEAN UP RESOURCES THE DELETED OBJECT OWNED.

Finalizers are the keys to tell Kubernetes API that, there are few resources to be deleted or taken care of before this particular resource is deleted.