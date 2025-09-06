# Knowledge Development Container

このディレクトリには、Dev Container設定が含まれています。

## 概要

Dev Containerは、Dockerコンテナ内で標準化された開発環境を提供し、以下のメリットがあります：

- **環境の一貫性**: チーム全体で同じ開発環境を共有
- **依存関係の管理**: 必要なツールやライブラリを自動的にインストール
- **セキュリティ**: 厳格なファイアウォール設定によるネットワーク制限
- **即座に開発開始**: リポジトリをクローンするだけで開発環境が整う

### 設定ファイル詳細

- `devcontainer.json`: Dev Container全体設定
- `Dockerfile`: コンテナ環境構築
- `init-firewall.sh`: ファイアウォール設定スクリプト

## 開発開始手順

### 1. VS Codeでの開発
1. VS CodeでプロジェクトフォルダーをOpen
2. "Reopen in Container"を選択
3. 初回はコンテナビルド（数分かかります）
4. 自動的にファイアウォール設定が適用される

### 2. GoLandでの開発
1. GoLandでプロジェクトフォルダーをOpen
2. Services(⌘8) で Show Dev Containers を選択
  - Rebuild か Start を選択

## 技術スタック

### 開発環境
- **Node.js**: 20 (Next.js フロントエンド用)
- **Go**: 1.24.5 (バックエンド用)
- **PostgreSQL Client**: データベース接続用
- **Claude Code**: AI支援開発ツール

### VS Code拡張機能
- **ESLint**: JavaScriptコード品質チェック
- **Prettier**: コードフォーマッター
- **Go**: Go言語サポート
- **GraphQL**: GraphQL構文サポート
- **GitLens**: Git履歴可視化

### セキュリティ機能
- **iptables/ipset**: 厳格なファイアウォール設定
- **許可ドメイン**: GitHub、npm、Anthropic API、Go modulesのみアクセス許可
- **非rootユーザー**: nodeユーザーでの安全な実行

## 開発ワークフロー

### ローカルでの作業

- (TODO)

### Dev Container での作業

- Claude Codeを利用したAI支援開発
- E2Eテスト実行
  - (TODO) E2E テストでは playwright.config.ts をベタガキで書き換えている
    - baseURL を http://host.docker.internal:3000
    - launchOptions を設定
  - (TODO) dev container postCreateCommand の playwright インストールがうまくいってないかも？

```bash
# E2Eテスト実行（Docker-in-Docker）
make dev-start-ci-devcontainer # CI環境起動
make run-e2e-only              # E2Eテスト実行

# Docker環境管理
make ci-devcontainer-compose-ps   # 実行中コンテナ確認
make health-check-ci-devcontainer # サービス状態確認
make ci-devcontainer-compose-stop # CI環境停止
make ci-devcontainer-compose-down # CI環境クリーンアップ
```

## ネットワーク設定

### 許可されているドメイン
- **GitHub**: リポジトリアクセス
- **npmjs.org**: Node.jsパッケージ
- **golang.org**: Go modules
- **anthropic.com**: Claude API
- **ローカルホスト**: Docker内部通信

### ブロックされるアクセス
- 上記以外の全外部通信
- example.comなどのテスト用ドメイン

## 更新・メンテナンス

### ツールバージョン更新
- (TODO)

### セキュリティ設定追加
- .devcontainer/init-firewall.sh に新しいドメインを追加
- domain listに追加

## トラブルシューティング

### よくある問題

**1. Go modulesのダウンロードが失敗する**
```bash
# ファイアウォール設定を確認
sudo iptables -L -n
# proxy.golang.org, sum.golang.orgが許可されているか確認
```

**2. npm installが失敗する**
```bash
# registry.npmjs.orgへの接続を確認
curl https://registry.npmjs.org
```

## 参考資料

- [VS Code Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers)
- [Docker Compose](https://docs.docker.com/compose/)
- [iptables Tutorial](https://netfilter.org/documentation/)
