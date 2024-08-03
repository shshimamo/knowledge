#
# EKS Cluster
#
.PHONY: create-cluster
create-cluster:
	eksctl create cluster -f eksctl-config.yaml

.PHONY: delete-cluster
delete-cluster:
	eksctl delete cluster -f eksctl-config.yaml --wait --disable-nodegroup-eviction

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

