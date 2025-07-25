# E2Eテスト

このディレクトリには、KnowledgeアプリケーションのE2E（エンドツーエンド）テストが含まれています。

## 前提条件

E2Eテストを実行する前に、以下のサービスが起動している必要があります：

### 1. バックエンドサービス

#### Backend Main (GraphQL API + DB)
```bash
cd backend/main
make database-create-develop
make migration-develop
docker-compose up
```
→ `localhost:8080` で起動

#### Backend Auth (認証API)
```bash
cd backend/auth
make database-create-develop
make migration-develop
docker-compose up
```
→ `localhost:80` で起動

### 2. フロントエンド
```bash
cd frontend/main
npm run dev
```
→ `localhost:3000` で起動

## テストの実行

### 基本的なテスト実行
```bash
# ヘッドレスモードでテスト実行
npm run test:e2e

# ブラウザを表示してテスト実行
npm run test:e2e:headed

# Playwright UIモードでテスト実行
npm run test:e2e:ui
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

## テストケース

### サインアップ機能
- ✅ サインアップフォームの表示確認
- ✅ 正常なサインアップ処理
- ✅ バリデーションエラーの表示
- ✅ 無効なメールアドレスのエラー処理
- ✅ 既存ユーザーの重複エラー処理

### ナレッジ機能
- ✅ ナレッジ作成フォームの表示確認
- ✅ プライベートナレッジの作成
- ✅ パブリックナレッジの作成
- ✅ ナレッジ一覧での表示確認
- ✅ ナレッジ作成時のバリデーション
- ✅ ナレッジの編集機能
- ✅ サインアップからナレッジ作成までの一連の流れ

## 設定ファイル

- `playwright.config.ts` - Playwright設定
  - テストディレクトリ: `./tests/e2e`
  - ベースURL: `http://localhost:3000`
  - ブラウザ: Chromium
  - 開発サーバー自動起動設定

## トラブルシューティング

### よくある問題

1. **テストが失敗する場合**
   - 全てのバックエンドサービスが起動していることを確認
   - フロントエンドが `localhost:3000` で起動していることを確認
   - データベースのマイグレーションが完了していることを確認

2. **ポートが使用されている場合**
   ```bash
   # ポート使用状況確認
   lsof -i :3000
   lsof -i :8080
   lsof -i :80
   
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

GitHub Actionsでの自動テスト実行も想定しており、CI環境では以下のように実行されます：

```bash
npm run test:e2e -- --reporter=github
```