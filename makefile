.PHONY: easy_setup

easydeploy:
	rm -rf certs
	git pull

	cp /home/certs .

	(cd v1/ && CGO_ENABLED=0 GOOS=linux make build)
	(cd v2/ && CGO_ENABLED=0 GOOS=linux make build)

	docker compose up