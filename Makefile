setup:
	docker-compose build
	docker-compose run --rm restapi ash -c "sql-migrate up && go run db/create_database.go"

db.create:
	docker-compose run --rm restapi go run db/create_database.go

db.migrate:
	docker-compose run --rm restapi sql-migrate up

db.seed:
	docker-compose run --rm restapi go run db/seed/seeder.go

start:
	docker-compose up

end:
	docker-compose down

entry-server-container:
	docker-compose exec restapi ash

entry-db-container:
	docker-compose exec db bash