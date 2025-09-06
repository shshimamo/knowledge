# Knowledge Development Container

このディレクトリには、knowledgeプロジェクトのDev Container設定が含まれています。

## 概要

Dev Containerは、Dockerコンテナ内で標準化された開発環境を提供し、以下のメリットがあります：

- **環境の一貫性**: チーム全体で同じ開発環境を共有
- **依存関係の管理**: 必要なツールやライブラリを自動的にインストール
- **セキュリティ**: 厳格なファイアウォール設定によるネットワーク制限
- **即座に開発開始**: リポジトリをクローンするだけで開発環境が整う

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

## ハイブリッド開発ワークフロー

### GoLand（ローカル）での作業
```
- Go コード編集・リファクタリング
- デバッグ・プロファイリング  
- ファイル管理・Git操作
- プロジェクト構造把握
```

### Dev Container での作業
```bash
# 基本的な開発作業
cd frontend/main && npm run dev  # フロントエンド開発サーバー
cd backend/main && go run .      # バックエンド実行
cd backend/main && go mod tidy   # Go modules管理

# E2Eテスト実行（Docker-in-Docker）
make setup-ci     # CI環境セットアップ
make start-ci     # 全サービス起動
make test-e2e-ci  # E2Eテスト実行

# または個別実行
cd frontend/main && npm run test:e2e

# Docker環境管理
docker ps                    # 実行中コンテナ確認
make check-ci-services      # サービス状態確認
make stop-ci                # CI環境停止
make clean-ci               # CI環境クリーンアップ
```

## 開発開始手順

### 1. VS Codeでの開発（推奨）
1. VS CodeでプロジェクトフォルダーをOpen
2. "Reopen in Container"を選択
3. 初回はコンテナビルド（数分かかります）
4. 自動的にファイアウォール設定が適用される

### 2. ハイブリッド開発
```bash
# Dev Container起動
make devcontainer-up

# GoLandでローカル開発
# GoLand → Settings → Go → GOROOT を Dev Container内のGoに設定
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

**3. PostgreSQL接続エラー**
```bash
# ローカルホスト通信が許可されているか確認
# docker-compose.devcontainer.ymlでDBサービスを有効化
```

### 設定ファイル詳細

- `devcontainer.json`: Dev Container全体設定
- `Dockerfile`: コンテナ環境構築
- `init-firewall.sh`: ファイアウォール設定スクリプト

## 更新・メンテナンス

### ツールバージョン更新
```bash
# Go version更新
vim .devcontainer/devcontainer.json
# GO_VERSION を更新

# コンテナ再ビルド
docker-compose -f docker-compose.devcontainer.yml build --no-cache
```

### セキュリティ設定追加
```bash
# 新しい許可ドメイン追加
vim .devcontainer/init-firewall.sh
# domain listに追加
```

## 参考資料

- [VS Code Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers)
- [Docker Compose](https://docs.docker.com/compose/)
- [iptables Tutorial](https://netfilter.org/documentation/)