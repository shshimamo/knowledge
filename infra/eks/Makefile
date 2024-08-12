AWS_ACCOUNT_ID = $(shell aws sts get-caller-identity --query Account --output text)

define check_var
  @if [ -z "$(1)" ]; then \
    echo "$(2) is not set"; \
    exit 1; \
  fi
endef

#
# EKS Cluster
#
.PHONY: create-cluster
create-cluster:
	sed "s/__AWS_ACCOUNT_ID__/${AWS_ACCOUNT_ID}/g" ./eksctl-config.yaml | eksctl create cluster -f -

.PHONY: dry-run-create-cluster
dry-run-create-cluster:
	sed "s/__AWS_ACCOUNT_ID__/${AWS_ACCOUNT_ID}/g" ./eksctl-config.yaml | eksctl create cluster --dry-run -f -

.PHONY: delete-cluster
delete-cluster:
	eksctl delete cluster -f eksctl-config.yaml --wait --disable-nodegroup-eviction

#
# Access Entry
#
.PHONY: create-accessentry
create-accessentry:
	eksctl create accessentry -f eksctl-config.yaml

.PHONY: delete-accessentry
delete-accessentry:
	eksctl delete accessentry --cluster KnowledgeCluster --principal-arn arn:aws:iam::$${AWS_ACCOUNT_ID}:role/github-oidc

#
# Namespace
#
.PHONY: create-namespace-app-ns
create-namespace-app-ns:
	kubectl create namespace app-ns

.PHONY: delete-namespace-app-ns
delete-namespace-app-ns:
	kubectl delete namespace app-ns

#
# Secret, ConfigMap
#
.PHONY: create-db-secret
create-db-secret:
	kubectl create secret generic db-credentials --from-literal=DB_PASSWORD=${DB_PASSWORD} --namespace=app-ns

.PHONY: create-db-configmap
create-db-configmap:
	kubectl create configmap db-config --from-literal=DB_HOST=${DB_HOST} --from-literal=DB_PORT=5432 --from-literal=DB_USER=postgres --namespace=app-ns

.PHONY: create-application-configmap
create-application-configmap:
	kubectl create configmap application-config --from-literal=APP_ENV=${APP_ENV} --namespace=app-ns

.PHONY: check-db-secret
check-db-secret:
	kubectl get secret --namespace=app-ns

.PHONY: check-db-configmap
check-db-configmap:
	kubectl get configmap db-config -o yaml --namespace=app-ns

.PHONY: check-application-configmap
check-application-configmap:
	kubectl get configmap application-config -o yaml --namespace=app-ns

#
# kubeconfig
#
.PHONY: update-kubeconfig
update-kubeconfig:
	aws eks --region ap-northeast-1 update-kubeconfig --name KnowledgeCluster

