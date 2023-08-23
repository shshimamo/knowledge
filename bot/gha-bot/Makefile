define check_var
  @if [ -z "$(1)" ]; then \
    echo "$(2) is not set"; \
    exit 1; \
  fi
endef

.PHONY: deploy
deploy:
	$(call check_var,${SLACK_SIGNING_SECRET},SLACK_SIGNING_SECRET)
	$(call check_var,${SLACK_BOT_TOKEN},SLACK_BOT_TOKEN)
	$(call check_var,${GITHUB_TOKEN},GITHUB_TOKEN)
	npx serverless deploy

.PHONY: remove
remove:
	npx serverless remove