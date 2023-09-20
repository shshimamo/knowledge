# sample/grpc_hello/grpc_client, grpc_server 概要

* grpc_server: サンプルgRPCサーバー
* grpc_client: サンプルgRPCサーバーのクライアント
   * /grpc にアクセスするとgRPCサーバーを呼び出す

# 構築手順
1. EKS: Cluster 作成
   * infra/README.md 参照

2. EKS: Namespace(app-ns) 作成
   * infra/README.md 参照

3. EKS: Service 作成

```sh
$ cd sample/grpc_hello/grpc_server
$ make apply-service
```

```sh
$ cd sample/grpc_hello/grpc_client
$ make apply-service
```

4. EKS: Ingress 作成
   * infra/README.md 参照
   * infra/eks/ingress の apply-ingress は不要

```sh
$ cd sample/grpc_hello/grpc_client
$ make apply-ingress
```

5. IAM: GitHub Actions 用のIDプロバイダ、ロール
    * infra/README.md 参照

6. EKS: RBAC
    * infra/README.md 参照

7. ECR
    * sample-grpc-client リポジトリ作成
   * sample-grpc-server リポジトリ作成

8. デプロイ
    * GHA の Build Push | Sample を実行
    * GHA の Deploy | Sample を実行

9. Route53
    * ALBへのエイリアスを作成(sample-grpc-client.shshimamo.com)
