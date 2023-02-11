up:
	docker compose up -d

sh :
	docker compose exec app sh

go_test:
	docker compose exec app /bin/sh -c 'go test ./...'
