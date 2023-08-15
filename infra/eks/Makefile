#
# EKS Cluster
#
.PHONY: create-cluster
create-cluster:
	eksctl create cluster -f eksctl-config.yaml

.PHONY: delete-cluster
delete-cluster:
	eksctl delete cluster -f eksctl-config.yaml

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
	kubectl create configmap db-config --from-literal=DB_HOST=${DB_HOST} --from-literal=DB_PORT=5432 --from-literal=DB_NAME=todo-app-next-go --from-literal=DB_USER=postgres --namespace=app-ns

#
# kubeconfig
#
.PHONY: update-kubeconfig
update-kubeconfig:
	aws eks --region ap-northeast-1 update-kubeconfig --name todo-cluster

