setup:
	docker-compose build
	docker-compose run --rm graphql sql-migrate up

start:
	docker-compose up

end:
	docker-compose down

entry-server-container:
	docker-compose exec graphql ash

entry-db-container:
	docker-compose exec db bash