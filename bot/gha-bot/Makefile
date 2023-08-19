.PHONY: deploy
deploy:
	@if [ -z "$${SLACK_SIGNING_SECRET}" ]; then \
		echo "SLACK_SIGNING_SECRET is not set"; \
		exit 1; \
	fi
	@if [ -z "$${SLACK_BOT_TOKEN}" ]; then \
		echo "SLACK_BOT_TOKEN is not set"; \
		exit 1; \
	fi
	@if [ -z "$${GITHUB_TOKEN}" ]; then \
		echo "GITHUB_TOKEN is not set"; \
		exit 1; \
	fi
	npx serverless deploy

.PHONY: remove
remove:
	npx serverless remove