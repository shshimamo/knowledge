.PHONY: apply-role
apply-role:
	kubectl apply -f role.yml -n app-ns

.PHONY: check-clusterroles
check-clusterroles:
	kubectl describe clusterroles role-for-github-actions

.PHONY: check-rolebindings
check-rolebindings:
	kubectl describe rolebindings -n app-ns;

.PHONY: create-iamidentitymapping
create-iamidentitymapping:
	export AWS_ACCOUNT_ID=$$(aws sts get-caller-identity --query Account --output text); \
	eksctl create iamidentitymapping \
	--cluster knowledge-cluster \
	--region=ap-northeast-1 \
    --arn arn:aws:iam::${AWS_ACCOUNT_ID}:role/github-oidc \
    --username github-actions  \
    --group github-actions \
    --no-duplicate-arns

.PHONY: check-iamidentitymapping
check-iamidentitymapping:
	eksctl get iamidentitymapping --cluster knowledge-cluster --region=ap-northeast-1