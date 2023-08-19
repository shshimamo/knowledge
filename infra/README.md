# 環境構築

## EKSクラスター作成
```sh
$ cd infra/eks
$ make create-cluster
```

## RDS作成
```sh
# eksctl で作成した VPC など利用
$ cd infra/rds
$ DB_PASSWORD=xxx make rds-cdk-deploy
```

## EKS: Namespace, ConfigMap, Secret 作成
```sh
$ cd infra/eks
$ export DB_PASSWORD=password
$ export DB_HOST=host

# Namespace 作成
$ make create-namespace-app-ns
# Secret 作成
$ make create-db-secret
# ConfigMap 作成
$ make create-db-configmap
```

## EKS: Service 作成
```sh
$ cd backend/main/k8s
$ make apply-service

$ cd backend/auth/k8s
$ make apply-service

$ cd frontend/main/k8s
$ make apply-service
```

## EKS: Ingress 作成
```sh
$ cd infra/eks/ingress

# OIDC Provider 作成
$ make create-oidc-provider

# Policy 作成
$ make create-ingress-controller-policy

# IAM Role と Service Account 作成
$ make create-iamserviceaccount

# Ingress Controller インストール
$ make install-ingress-controller

# Ingress 作成
$ make apply-ingress
```

## IAM: GA用のIDプロバイダ、ロール

* TODO: 自動化

## EKS: Role と IAM プリンシパル
```sh
$ cd infra/eks/role
# ClusterRole, RoleBinding 作成
$ make apply-role
# IAM プリンシパル 作成
$ make create-iamidentitymapping
```

## ECR

手動作成
```
knowledge-backend-auth
knowledge-backend-auth-migration
knowledge-backend-main
knowledge-backend-main-migration
knowledge-frontend-main
```

# デプロイ

## バックエンドデプロイ(Main)
* GHA の Build Push | Backend Main を実行
* GHA の Deploy | Backend Main を実行

## バックエンドデプロイ(Auth)
* GHA の Build Push | Backend Auth を実行
* GHA の Deploy | Backend Auth を実行

## フロントエンドデプロイ
* GHA の Build Push | Frontend Main を実行
* GHA の Deploy | Frontend Main を実行

## マイグレーション(Main)
* GHA の Build Push | Backend Main Migration を実行
* GHA の Migrate | Backend Main を実行

## マイグレーション(Auth)
* GHA の Build Push | Backend Auth Migration を実行
* GHA の Migrate | Backend Auth を実行

# クリーンアップ

## Namespace 削除
```sh
$ cd infra/eks
$ make delete-namespace-app-ns
```

## RDS 削除
```sh
$ cd infra/rds
$ make rds-cdk-destroy
```

## EKS 削除
```sh
$ cd infra/eks
$ make delete-cluster
```