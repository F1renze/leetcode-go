
update: ## update Readme
	@go run main.go

count: ## count total solutions
	@ls -l solutions/|wc -l

.DEFAULT_GOAL :=
help: ## 显示帮助
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
