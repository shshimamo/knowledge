.PHONY: apply-role
apply-role:
	kubectl apply -f role.yml -n app-ns

.PHONY: check-clusterroles
check-clusterroles:
	kubectl describe clusterroles role-for-github-actions

.PHONY: check-clusterroles2
check-clusterroles2:
	kubectl get clusterrole role-for-github-actions -o yaml

.PHONY: check-rolebindings
check-rolebindings:
	kubectl describe rolebindings -n app-ns;

.PHONY: search-rolebindings
search-rolebindings:
	kubectl get rolebindings,clusterrolebindings --all-namespaces | grep github-actions

.PHONY: create-iamidentitymapping
create-iamidentitymapping:
	AWS_ACCOUNT_ID=$(shell aws sts get-caller-identity --query Account --output text) ;\
	echo "AWS_ACCOUNT_ID: $$AWS_ACCOUNT_ID" ;\
	eksctl create iamidentitymapping \
	--cluster knowledge-cluster \
	--region=ap-northeast-1 \
    --arn arn:aws:iam::$$AWS_ACCOUNT_ID:role/github-oidc \
    --username github-actions  \
    --group github-actions \
    --no-duplicate-arns

.PHONY: delete-iamidentitymapping
delete-iamidentitymapping:
	AWS_ACCOUNT_ID=$(shell aws sts get-caller-identity --query Account --output text) ;\
	echo "AWS_ACCOUNT_ID: $$AWS_ACCOUNT_ID" ;\
	eksctl delete iamidentitymapping --cluster knowledge-cluster --arn arn:aws:iam::$$AWS_ACCOUNT_ID:role/github-oidc


.PHONY: check-iamidentitymapping
check-iamidentitymapping:
	eksctl get iamidentitymapping --cluster knowledge-cluster --region=ap-northeast-1