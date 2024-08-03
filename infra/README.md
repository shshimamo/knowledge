# 環境構築

## EKSクラスター作成
```sh
$ cd infra/eks
$ make create-cluster
```

## RDS作成
* 注意: デフォルトパラメータグループの rds.force_ssl が 1 になっているとEKSから接続エラーになるのでアプリ対応するまでは 0 にしておく
```sh
# eksctl で作成した VPC など利用
$ cd infra/rds
$ DB_PASSWORD=xxx make rds-cdk-deploy
```

## EKS: Namespace, ConfigMap, Secret 作成
```sh
$ cd infra/eks
$ export DB_PASSWORD=xxx
# rds-cdk-deploy で作成した RDS のエンドポイント
$ export DB_HOST=xxx.ap-northeast-1.rds.amazonaws.com

# Namespace 作成
$ make create-namespace-app-ns
# Secret 作成
$ make create-db-secret
# ConfigMap 作成
$ make create-db-configmap
$ export APP_ENV=production
$ make create-application-configmap
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

# EKS に OIDC Provider 作成
$ make create-oidc-provider

# IAM Policy 作成
$ make create-ingress-controller-policy

# IAM Role と Service Account 作成
$ make create-iamserviceaccount

# Ingress Controller インストール
$ make install-ingress-controller

# Ingress 作成(少しIngress Controllerインストールを待つ必要があるかも)
$ make apply-ingress
```

## IAM: GitHub Actions 用のIDプロバイダ、ロール

* TODO: 自動化

## EKS: RBAC
```sh
$ cd infra/eks/role
# ClusterRole, RoleBinding 作成
$ make apply-role
# GHA が使う IAM ロールをk8s内のユーザー、グループに紐付ける
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
