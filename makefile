.PHONY: easy_setup

easydeploy:
	git pull

	(cd v1/ && CGO_ENABLED=0 GOOS=linux make build)
	(cd v2/ && CGO_ENABLED=0 GOOS=linux make build)

	cp -rf /home/admin/frontend ./frontend

	docker compose up --build