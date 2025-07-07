.PHONY: help generate_swagger easydeploy
help:
	@echo "Available commands:"
	@echo "  generate_swagger  - Generate swagger documentation for v1 and v2"
	@echo "  easydeploy	   - Pull changes, build v1 and v2, copy frontend and deploy with docker"

generate_swagger:
	(cd v1/ && swag init -g cmd/main.go -d ./ --parseDependency --parseInternal)
	(cd v2/ && swag init -g cmd/main.go -d ./ --parseDependency --parseInternal)

easydeploy:
	go install github.com/swaggo/swag/cmd/swag@latest

	make generate_swagger
	git fetch --all
	git reset --hard origin/$(shell git rev-parse --abbrev-ref HEAD)

	(cd v1/ && CGO_ENABLED=0 GOOS=linux make build)
	(cd v2/ && CGO_ENABLED=0 GOOS=linux make build)

	cp -rf /home/admin/frontend ./frontend

	docker compose up --build
