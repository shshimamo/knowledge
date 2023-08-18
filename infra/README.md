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

## Namespace, ConfigMap, Secret 作成
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

## Service 作成
```sh
$ make apply-backend-service
```

## Ingress 作成
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

## Role と IAM プリンシパル
```sh
$ cd infra/eks/role
# ClusterRole, RoleBinding 作成
$ make apply-role
# IAM プリンシパル 作成
$ make create-iamidentitymapping
```

## マイグレーション
* GHA の Push Image Migration を実行
* GHA の Exec Migration Job を実行

# デプロイ

## バックエンドデプロイ
* GHA の Push Image Backend を実行
* GHA の Apply Backend Deployment を実行

## フロントエンドデプロイ
* GHA の Push Image Frontend を実行
* GHA の Apply Frontend Deployment を実行

## マイグレーション
* GHA の Push Image Migration を実行
* GHA の Exec Migration Job を実行

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