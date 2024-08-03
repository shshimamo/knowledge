#
# OIDC Provider
#
.PHONY: create-oidc-provider
create-oidc-provider:
	eksctl utils associate-iam-oidc-provider --cluster KnowledgeCluster --approve

#
# Policy
#
.PHONY: create-ingress-controller-policy
create-ingress-controller-policy:
	aws iam create-policy \
		--policy-name AWSLoadBalancerControllerIAMPolicy \
		--policy-document file://iam_policy.json

#
# IAM Role and Service Account
#
.PHONY: create-iamserviceaccount
create-iamserviceaccount:
	export AWS_ACCOUNT_ID=$$(aws sts get-caller-identity --query Account --output text); \
	eksctl create iamserviceaccount \
		--cluster=KnowledgeCluster \
		--namespace=kube-system \
		--name=aws-load-balancer-controller \
		--role-name AmazonEKSLoadBalancerControllerRole \
		--attach-policy-arn=arn:aws:iam::$${AWS_ACCOUNT_ID}:policy/AWSLoadBalancerControllerIAMPolicy \
		--approve

.PHONY: check-create-iamserviceaccount--role
check-create-iamserviceaccount--role:
	aws iam get-role --role-name AmazonEKSLoadBalancerControllerRole

.PHONY: check-create-iamserviceaccount--sa
check-create-iamserviceaccount--sa:
	kubectl get sa -n kube-system aws-load-balancer-controller -o yaml

#
# Ingress Controller
#
.PHONY: install-ingress-controller
install-ingress-controller:
	helm repo add eks https://aws.github.io/eks-charts; \
	helm repo update; \
	helm install aws-load-balancer-controller eks/aws-load-balancer-controller \
		-n kube-system \
		--set clusterName=KnowledgeCluster \
		--set serviceAccount.create=false \
		--set serviceAccount.name=aws-load-balancer-controller

.PHONY: check-install-ingress-controller
check-install-ingress-controller:
	kubectl get pod -n kube-system -l app.kubernetes.io/name=aws-load-balancer-controller

#
# Ingress Resource
#
.PHONY: apply-ingress
apply-ingress:
	kubectl apply -f ingress.yaml -n app-ns
