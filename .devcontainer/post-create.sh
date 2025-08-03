#!/bin/bash

# Dev Container のセットアップ後に実行されるスクリプト

echo "🚀 Dev Container のセットアップを開始します..."

# Git の設定確認（コンテナ内での開発用）
echo "📝 Git 設定の確認..."
git config --global --get user.name || git config --global user.name "Claude Code User"
git config --global --get user.email || git config --global user.email "user@example.com"

# Go モジュールの初期化とダウンロード
echo "🔧 Go モジュールのセットアップ..."
if [ -d "backend/main" ]; then
    cd backend/main && go mod download
    cd /workspaces/knowledge
fi

if [ -d "backend/auth" ]; then
    cd backend/auth && go mod download
    cd /workspaces/knowledge
fi

# Node.js 依存関係のインストール
echo "📦 Node.js 依存関係のインストール..."
if [ -d "frontend/main" ]; then
    cd frontend/main && npm install
    cd /workspaces/knowledge
fi

# Playwright ブラウザのインストール
echo "🎭 Playwright ブラウザのインストール..."
if [ -d "frontend/main" ]; then
    cd frontend/main && npx playwright install
    cd /workspaces/knowledge
fi

# Go ツールのインストール
echo "🛠️ Go 開発ツールのインストール..."
go install golang.org/x/tools/gopls@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/go-delve/delve/cmd/dlv@latest

# Claude Code の設定確認
echo "🤖 Claude Code の設定確認..."
if command -v claude &> /dev/null; then
    echo "✅ Claude Code がインストールされています"
    echo "💡 初回利用時は 'claude auth' でAPI認証を行ってください"
else
    echo "⚠️ Claude Code のインストールに失敗しました"
fi

echo "🎉 Dev Container のセットアップが完了しました！"
echo ""
echo "💡 利用可能なサービス:"
echo "   - フロントエンド: http://localhost:3000"
echo "   - バックエンド API: http://localhost:8000"
echo "   - GraphQL: http://localhost:8080"
echo "   - 認証 API: http://localhost:8081"
echo "   - PostgreSQL: localhost:5432"
echo ""
echo "🏃‍♂️ 開発を開始するには:"
echo "   cd frontend/main && npm run dev    # フロントエンド起動"
echo "   cd backend/main && make run        # バックエンド起動"
echo "   cd backend/auth && make run        # 認証サービス起動"