.PHONY: easy_setup

easy_setup:
	git pull

	(cd v1/ && CGO_ENABLED=0 GOOS=linux make build)
	(cd v2/ && CGO_ENABLED=0 GOOS=linux make build)

	docker compose up