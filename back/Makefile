.PHONY: up down build db

up:
	docker-compose up -d

down :
	docker-compose down

build:
	docker-compose build

db:
	docker exec -it back-database-1 bash -c "psql -U root -d goNine"