setup:
	docker-compose build
	docker-compose run --rm restapi go run db/create_database.go
	docker-compose run --rm restapi sql-migrate up

start:
	docker-compose up

end:
	docker-compose down

entry-server-container:
	docker-compose exec restapi ash

entry-db-container:
	docker-compose exec db bash