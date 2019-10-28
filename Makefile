
update: ## 更新 Readme
	@go run main.go

.DEFAULT_GOAL :=
help: ## 显示帮助
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
