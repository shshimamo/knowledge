#
# CI環境関連のコマンド
#

# CI環境でのセットアップ（データベース作成とマイグレーション）
setup-ci:
	@echo "CI環境をセットアップ中..."
	docker compose -f docker-compose.ci.yml up -d db
	@echo "PostgreSQLの起動を待機中..."
	sleep 10
	@echo "データベースとマイグレーションを実行中..."
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=test make -C backend/main database-create
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=test make -C backend/main migration
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=test make -C backend/auth database-create
	DB_HOST=localhost DB_PASSWORD=password APP_ENV=test make -C backend/auth migration
	@echo "CI環境のセットアップが完了しました"

# CI環境での全サービス起動
start-ci:
	docker compose -f docker-compose.ci.yml up -d

# CI環境でのE2Eテスト実行（自動でstart-ciも実行）
test-e2e-ci: setup-ci start-ci
	sleep 30
	cd frontend/main && CI=true npm run test:e2e

# CI環境停止
stop-ci:
	@echo "CI環境を停止中..."
	docker compose -f docker-compose.ci.yml stop

# CI環境のクリーンアップ
clean-ci:
	docker compose -f docker-compose.ci.yml down -v
	docker system prune -f

clean-and-rmi-ci:
	docker compose -f docker-compose.ci.yml down -v --rmi all || true
	docker system prune -f

#
# ローカル開発用コマンド
#

# CI環境が既に起動しているかチェック
check-ci-running:
	@if docker compose -f docker-compose.ci.yml ps --services --filter "status=running" | grep -q .; then \
		echo "⚠️  CI環境は既に起動しています"; \
		echo "サービス状態: make check-ci-services"; \
		echo "停止する場合: make stop-ci"; \
		exit 1; \
	fi

# CI環境を起動してバックグラウンドで実行（開発用）
dev-start-ci: check-ci-running setup-ci start-ci
	@echo "==============================================="
	@echo "✅ CI環境が起動しました！"
	@echo "フロントエンド: http://localhost:3000"
	@echo "バックエンドAPI: http://localhost:8080"
	@echo "認証API: http://localhost:8081"
	@echo "PostgreSQL: localhost:5432"
	@echo ""
	@echo "📋 利用可能なコマンド:"
	@echo "  サービス確認: make check-ci-services"
	@echo "  E2Eテスト実行: make run-e2e-only"
	@echo "  環境停止: make stop-ci"
	@echo "==============================================="

# CI環境を強制的に再起動（既存環境を停止してから起動）
dev-restart-ci:
	@echo "CI環境を再起動中..."
	make stop-ci || true
	make clean-ci || true
	make dev-start-ci
	@echo "サービスの起動を待機中（30秒）..."
	sleep 30
	@echo "サービス状態を確認します..."
	make check-ci-services

# CI環境を強制的に再起動（既存環境を停止してから起動）
# 使用例: make dev-restart-and-rmi-ci SERVICES="frontend"
# 使用例: make dev-restart-and-rmi-ci SERVICES="frontend backend-main backend-auth"
# 使用例: make dev-restart-and-rmi-ci （全サービス）
dev-restart-and-rmi-ci:
	@echo "CI環境を再起動中..."
	make stop-ci || true
	@echo "SERVICES: ${SERVICES}"
	@if [ "${SERVICES}" ]; then \
		echo "指定されたサービスイメージを削除"; \
		for service in ${SERVICES}; do \
			IMAGE_ID=$$(docker compose -f docker-compose.ci.yml images -q $${service} 2>/dev/null); \
			if [ "$${IMAGE_ID}" ]; then \
				echo "コンテナを削除中: $${service}"; \
				docker compose -f docker-compose.ci.yml rm -f $${service} || true; \
				echo "イメージを削除中: $${service}/$${IMAGE_ID}"; \
				docker rmi $${IMAGE_ID} || true; \
			fi; \
		done; \
	else \
		echo "全てのサービスイメージを削除"; \
		make clean-and-rmi-ci || true; \
	fi
	make dev-start-ci
	@echo "サービスの起動を待機中（30秒）..."
	sleep 30
	@echo "サービス状態を確認します..."
	make check-ci-services

#
# E2Eテスト実行
#   参照: frontend/main/tests/e2e/README.md
#

# 起動中のCI環境でE2Eテストのみ実行
run-e2e-only:
	@echo "E2Eテストを実行中..."
	cd frontend/main && npm run test:e2e

# E2Eテストを実行（UIモード）
run-e2e-ui:
	@echo "E2EテストをUIモードで実行中..."
	cd frontend/main && npm run test:e2e:ui

# E2Eテストを実行（ヘッドモード）
run-e2e-headed:
	@echo "E2Eテストをヘッドモードで実行中..."
	cd frontend/main && npm run test:e2e:headed

#
# CI環境のサービス状態確認
#
check-ci-services:
	@echo "CI環境のサービス状態を確認中..."
	@docker compose -f docker-compose.ci.yml ps
	@echo ""
	@echo "サービスヘルスチェック:"
	@echo -n "PostgreSQL: "
	@if docker compose -f docker-compose.ci.yml exec -T db pg_isready -U postgres > /dev/null 2>&1; then echo "✅ OK"; else echo "❌ NG"; fi
	@echo -n "Backend Main: "
	@if curl -f -s http://localhost:8080 > /dev/null 2>&1; then echo "✅ OK"; else echo "❌ NG"; fi
	@echo -n "Backend Auth: "
	@if curl -f -s http://localhost:8081 > /dev/null 2>&1; then echo "✅ OK"; else echo "❌ NG"; fi
	@echo -n "Frontend: "
	@if curl -f -s http://localhost:3000 > /dev/null 2>&1; then echo "✅ OK"; else echo "❌ NG"; fi

# CI環境のサービス状態確認（詳細版）
check-ci-services-detail:
	@echo "CI環境の詳細状態を確認中..."
	@echo "======================================"
	@echo "Docker Compose Services:"
	@docker compose -f docker-compose.ci.yml ps
	@echo ""
	@echo "======================================"
	@echo "ポート使用状況:"
	@echo "Port 5432 (PostgreSQL):"
	@lsof -i :5432 || echo "  ポート5432は使用されていません"
	@echo "Port 8080 (Backend Main):"
	@lsof -i :8080 || echo "  ポート8080は使用されていません"
	@echo "Port 8081 (Backend Auth):"
	@lsof -i :8081 || echo "  ポート8081は使用されていません"
	@echo "Port 3000 (Frontend):"
	@lsof -i :3000 || echo "  ポート3000は使用されていません"
	@echo ""
	@echo "======================================"
	@echo "サービスログ（最新10行）:"
	@echo "Backend Main:"
	@docker compose -f docker-compose.ci.yml logs --tail=10 backend-main || echo "  ログを取得できませんでした"
	@echo ""
	@echo "Backend Auth:"
	@docker compose -f docker-compose.ci.yml logs --tail=10 backend-auth || echo "  ログを取得できませんでした"
	@echo ""
	@echo "Frontend:"
	@docker compose -f docker-compose.ci.yml logs --tail=10 frontend || echo "  ログを取得できませんでした"

#
# CI環境のログ表示
#
logs-ci:
	docker compose -f docker-compose.ci.yml logs -f

logs-ci-frontend:
	docker compose -f docker-compose.ci.yml logs -f frontend

logs-ci-backend-main:
	docker compose -f docker-compose.ci.yml logs -f backend-main

logs-ci-backend-auth:
	docker compose -f docker-compose.ci.yml logs -f backend-auth

logs-ci-db:
	docker compose -f docker-compose.ci.yml logs -f db
