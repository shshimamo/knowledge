.PHONY: rds-cdk-deploy
rds-cdk-deploy:
	CDK_VPC_ID=$(shell aws ec2 describe-vpcs --filter "Name=tag:Name,Values=eksctl-KnowledgeCluster-cluster/VPC" --query "Vpcs[*].VpcId" --output text) ;\
	CDK_EKS_NODE_SG_ID=$(shell aws ec2 describe-security-groups --query 'SecurityGroups[*].[GroupName,GroupId]' --output text | grep eksctl-KnowledgeCluster-nodegroup-EksctlNodegroups1 | head -n1 | awk '{print $$2}') ;\
	CDK_MY_IP=$(shell curl -s https://checkip.amazonaws.com) ;\
	CDK_DB_PASSWORD=$$DB_PASSWORD ;\
	echo "CDK_VPC_ID: $$CDK_VPC_ID" ;\
	echo "CDK_EKS_NODE_SG_ID: $$CDK_EKS_NODE_SG_ID" ;\
	echo "CDK_MY_IP: $$CDK_MY_IP" ;\
	echo "CDK_DB_PASSWORD: $$CDK_DB_PASSWORD" ;\
	CDK_VPC_ID=$$CDK_VPC_ID CDK_EKS_NODE_SG_ID=$$CDK_EKS_NODE_SG_ID CDK_MY_IP=$$CDK_MY_IP CDK_DB_PASSWORD=$$CDK_DB_PASSWORD npx cdk deploy

.PHONY: rds-cdk-destroy
rds-cdk-destroy:
	CDK_VPC_ID=$(shell aws ec2 describe-vpcs --filter "Name=tag:Name,Values=eksctl-KnowledgeCluster-cluster/VPC" --query "Vpcs[*].VpcId" --output text) ;\
	CDK_EKS_NODE_SG_ID=$(shell aws ec2 describe-security-groups --query 'SecurityGroups[*].[GroupName,GroupId]' --output text | grep eksctl-KnowledgeCluster-nodegroup-EksctlNodegroups1 | head -n1 | awk '{print $$2}') ;\
	CDK_MY_IP=$(shell curl -s https://checkip.amazonaws.com) ;\
	CDK_DB_PASSWORD=$$DB_PASSWORD ;\
	echo "CDK_VPC_ID: $$CDK_VPC_ID" ;\
	echo "CDK_EKS_NODE_SG_ID: $$CDK_EKS_NODE_SG_ID" ;\
	echo "CDK_MY_IP: $$CDK_MY_IP" ;\
	echo "CDK_DB_PASSWORD: $$CDK_DB_PASSWORD" ;\
	CDK_VPC_ID=$$CDK_VPC_ID CDK_EKS_NODE_SG_ID=$$CDK_EKS_NODE_SG_ID CDK_MY_IP=$$CDK_MY_IP CDK_DB_PASSWORD=$$CDK_DB_PASSWORD npx cdk destroy

# 初回デプロイのみ
.PHONY: rds-cdk-bootstrap
rds-cdk-bootstrap:
	CDK_VPC_ID=$(shell aws ec2 describe-vpcs --filter "Name=tag:Name,Values=eksctl-KnowledgeCluster-cluster/VPC" --query "Vpcs[*].VpcId" --output text) ;\
	CDK_EKS_NODE_SG_ID=$(shell aws ec2 describe-security-groups --query 'SecurityGroups[*].[GroupName,GroupId]' --output text | grep eksctl-KnowledgeCluster-nodegroup-EksctlNodegroups1 | head -n1 | awk '{print $$2}') ;\
	CDK_MY_IP=$(shell curl -s https://checkip.amazonaws.com) ;\
	CDK_DB_PASSWORD=$$DB_PASSWORD ;\
	echo "CDK_VPC_ID: $$CDK_VPC_ID" ;\
	echo "CDK_EKS_NODE_SG_ID: $$CDK_EKS_NODE_SG_ID" ;\
	echo "CDK_MY_IP: $$CDK_MY_IP" ;\
	echo "CDK_DB_PASSWORD: $$CDK_DB_PASSWORD" ;\
	CDK_VPC_ID=$$CDK_VPC_ID CDK_EKS_NODE_SG_ID=$$CDK_EKS_NODE_SG_ID CDK_MY_IP=$$CDK_MY_IP CDK_DB_PASSWORD=$$CDK_DB_PASSWORD npx cdk bootstrap