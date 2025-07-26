# E2Eテスト

このディレクトリには、KnowledgeアプリケーションのE2E（エンドツーエンド）テストが含まれています。

## 実行手順

E2Eテストを実行する前に、以下のサービスが起動している必要があります：

```bash
# データベース作成、マイグレーション、サービス起動
make dev-start-ci

# サービス起動確認
make check-ci-services

# E2Eテスト実行
make run-e2e-only
```

frontendを修正した場合
```bash
# frontendイメージ再作成
make dev-restart-and-rmi-ci SERVICES="frontend"
```

## テストの実行

### 基本的なテスト実行
```bash
# ヘッドレスモードでテスト実行
npm run test:e2e

# makeコマンド
make run-e2e-only
```

```bash
# ブラウザを表示してテスト実行
npm run test:e2e:headed

# makeコマンド
make run-e2e-headed
```

```bash
# Playwright UIモードでテスト実行
npm run test:e2e:ui

# makeコマンド
make run-e2e-ui
```

### 特定のテストファイルのみ実行
```bash
# サインアップテストのみ実行
npx playwright test signup.spec.ts

# ナレッジ作成テストのみ実行
npx playwright test knowledge.spec.ts
```

## テスト構成

### テストファイル
- `signup.spec.ts` - サインアップ機能のテスト
- `knowledge.spec.ts` - ナレッジ作成・編集機能のテスト

### ヘルパーファイル
- `helpers/test-helper.ts` - 共通のテストユーティリティ関数

## 設定ファイル

- `playwright.config.ts` - Playwright設定
  - テストディレクトリ: `./tests/e2e`
  - ベースURL: `http://localhost:3000`
  - ブラウザ: Chromium
  - 開発サーバー自動起動設定

## トラブルシューティング

### よくある問題

1. **テストが失敗する場合**
   - 全てのサービスが起動していることを確認
     - `make check-ci-services`
     - `make check-ci-services-detail` (詳細版)
   - データベースのマイグレーションが完了していることを確認
   - CI環境のログを確認
     - `make logs-ci`
     - `make logs-ci-frontend`
     - `meke logs-ci-backend-main`
     - `make logs-ci-backend-auth`
     - `make logs-ci-db`

2. **ポートが使用されている場合**
   ```bash
   # ポート使用状況確認
   lsof -i :3000
   lsof -i :8080
   lsof -i :8081
   
   # プロセス終了
   kill -9 <PID>
   ```

3. **データベースエラーの場合**
   ```bash
   # データベース再作成
   cd backend/main
   make database-drop-develop
   make database-create-develop
   make migration-develop
   ```

## CI/CD

GitHub Actions で実行(e2e-test.yml)
