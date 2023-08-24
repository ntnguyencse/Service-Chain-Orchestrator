kubebuilder init --plugins go/v3 --domain automation.dcn.ssu.ac.kr --owner "Nguyen Thanh Nguyen" --project-name "sfc-at-edge" --repo "github.com/ntnguyencse/Service-Chain-Orchestrator"

kubebuilder create api --controller true --group sfc --version v1 --kind ServiceFunctionChain  --resource true
kubebuilder create api --controller true --group sfc --version v1 --kind ServiceLevelAgreement  --resource true
kubebuilder create api --controller true --group sfc --version v1 --kind Scheduler  --resource true
kubebuilder create api --controller true --group sfc --version v1 --kind SFCService  --resource true
kubebuilder create api --controller true --group sfc --version v1 --kind SFCDeployment  --resource true