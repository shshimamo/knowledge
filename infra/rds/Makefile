.PHONY: rds-cdk-deploy
rds-cdk-deploy:
	CDK_VPC_ID=$(shell aws ec2 describe-vpcs --filter "Name=tag:Name,Values=eksctl-knowledge-cluster-cluster/VPC" --query "Vpcs[*].VpcId" --output text) ;\
	CDK_EKS_NODE_SG_ID=$(shell aws ec2 describe-security-groups --query 'SecurityGroups[*].[GroupName,GroupId]' --output text | grep eksctl-knowledge-cluster-nodegroup-ng | head -n1 | awk '{print $$2}') ;\
	CDK_MY_IP=$(shell curl -s https://checkip.amazonaws.com) ;\
	CDK_DB_PASSWORD=$$DB_PASSWORD ;\
	echo "CDK_VPC_ID: $$CDK_VPC_ID" ;\
	echo "CDK_EKS_NODE_SG_ID: $$CDK_EKS_NODE_SG_ID" ;\
	echo "CDK_MY_IP: $$CDK_MY_IP" ;\
	echo "CDK_DB_PASSWORD: $$CDK_DB_PASSWORD" ;\
	CDK_VPC_ID=$$CDK_VPC_ID CDK_EKS_NODE_SG_ID=$$CDK_EKS_NODE_SG_ID CDK_MY_IP=$$CDK_MY_IP CDK_DB_PASSWORD=$$CDK_DB_PASSWORD npx cdk deploy

.PHONY: rds-cdk-destroy
rds-cdk-destroy:
	CDK_VPC_ID=$(shell aws ec2 describe-vpcs --filter "Name=tag:Name,Values=eksctl-knowledge-cluster-cluster/VPC" --query "Vpcs[*].VpcId" --output text) ;\
	CDK_EKS_NODE_SG_ID=$(shell aws ec2 describe-security-groups --query 'SecurityGroups[*].[GroupName,GroupId]' --output text | grep eksctl-knowledge-cluster-nodegroup-ng | head -n1 | awk '{print $$2}') ;\
	CDK_MY_IP=$(shell curl -s https://checkip.amazonaws.com) ;\
	CDK_DB_PASSWORD=$$DB_PASSWORD ;\
	echo "CDK_VPC_ID: $$CDK_VPC_ID" ;\
	echo "CDK_EKS_NODE_SG_ID: $$CDK_EKS_NODE_SG_ID" ;\
	echo "CDK_MY_IP: $$CDK_MY_IP" ;\
	echo "CDK_DB_PASSWORD: $$CDK_DB_PASSWORD" ;\
	CDK_VPC_ID=$$CDK_VPC_ID CDK_EKS_NODE_SG_ID=$$CDK_EKS_NODE_SG_ID CDK_MY_IP=$$CDK_MY_IP CDK_DB_PASSWORD=$$CDK_DB_PASSWORD npx cdk destroy

.PHONY: database-create
database-create:
	psql -h ${DB_HOST} -p 5432 -U postgres -c "CREATE DATABASE \"knowledge-main\";" ;\
	psql -h ${DB_HOST} -p 5432 -U postgres -c "CREATE DATABASE \"knowledge-auth\";"
