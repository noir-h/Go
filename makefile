up:
	docker-compose up -d
build:
	docker-compose build --no-cache --force-rm --platform amd64
stop:
	docker-compose stop
down:
	docker-compose down --remove-orphans
ps:
	docker-compose ps
logs:
	docker-compose logs   
go:
	docker compose exec app sh
run:
	docker compose exec app go run main.go
