.PHONY: apply-role
apply-role:
	kubectl apply -f role.yml -n app-ns

.PHONY: update-aws-auth
update-aws-auth:
	eksctl create iamidentitymapping \
	--cluster knowledge-cluster \
	--region=ap-northeast-1 \
    --arn arn:aws:iam::${AWS_ACCOUNT_ID}:role/github-oidc \
    --username github-actions  \
    --group github-actions \
    --no-duplicate-arns