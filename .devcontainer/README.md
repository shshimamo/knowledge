# Dev Container セットアップガイド

このプロジェクトでは VS Code の Dev Container 機能を使用して一貫した開発環境を提供しています。

## 必要な要件

- [Visual Studio Code](https://code.visualstudio.com/)
- [Dev Containers 拡張機能](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
- [Docker Desktop](https://www.docker.com/products/docker-desktop)

## セットアップ手順

1. このリポジトリをクローンします：
   ```bash
   git clone <repository-url>
   cd knowledge
   ```

2. VS Code でプロジェクトを開きます：
   ```bash
   code .
   ```

3. VS Code の右下に表示される「Reopen in Container」をクリックするか、
   コマンドパレット（Cmd/Ctrl + Shift + P）を開いて以下を実行：
   ```
   Dev Containers: Reopen in Container
   ```

4. 初回起動時は Docker イメージのビルドに数分かかります。
   完了後、自動的に開発環境がセットアップされます。

## 含まれている開発ツール

### 言語・ランタイム
- Go 1.24.5
- Node.js 18
- PostgreSQL クライアント

### 開発ツール
- Go 言語サーバー (gopls)
- golangci-lint
- Delve デバッガー
- Claude Code CLI
- GitHub CLI
- Docker & Docker Compose
- Playwright テストブラウザ

### VS Code 拡張機能
- Go 拡張機能
- TypeScript サポート
- Tailwind CSS IntelliSense
- Playwright テストランナー
- GraphQL サポート
- ESLint & Prettier
- Jest テストランナー

## サービス起動

Dev Container 内で以下のコマンドでサービスを起動できます：

### データベース
```bash
# データベースは自動的に起動されます
# 接続情報: localhost:5432, user: postgres, password: password
```

### バックエンドサービス
```bash
# メインAPI (GraphQL)
cd backend/main
make run

# 認証API
cd backend/auth
make run
```

### フロントエンド
```bash
cd frontend/main
npm run dev
```

## 利用可能なポート

- **3000**: フロントエンド (Next.js)
- **8000**: バックエンド メインAPI
- **8080**: GraphQL エンドポイント  
- **8081**: 認証API
- **5432**: PostgreSQL データベース

## Claude Code の使用

Dev Container 内で Claude Code を使用できます：

```bash
# Claude Code の起動
claude

# プロジェクト固有の設定は CLAUDE.md を参照
```

## トラブルシューティング

### コンテナが起動しない場合
1. Docker Desktop が起動していることを確認
2. メモリ不足の場合は Docker Desktop の設定でメモリを増やす
3. `Dev Containers: Rebuild Container` を実行

### データベースに接続できない場合
```bash
# データベースの状態確認
pg_isready -h db -p 5432 -U postgres

# マイグレーション実行
cd backend/main && make migrate-up
cd backend/auth && make migrate-up
```

### Node.js 依存関係の問題
```bash
cd frontend/main
rm -rf node_modules package-lock.json
npm install
```

## 開発ワークフロー

1. VS Code で Dev Container を開く
2. 各サービスを必要に応じて起動
3. ブラウザで http://localhost:3000 にアクセス
4. コードの変更は自動的にホストマシンと同期
5. テストとデバッグを Dev Container 内で実行

## 注意事項

- Dev Container 内でのファイル変更はホストマシンと自動同期されます
- データベースのデータは Docker ボリュームに永続化されます
- 初回セットアップ時のみ時間がかかります（2回目以降は高速）