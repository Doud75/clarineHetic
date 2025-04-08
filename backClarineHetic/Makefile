include .env
DOCKER_COMPOSE:=docker compose
UP:=up -d --build
DOWN:=stop
CONTAINER_DB_NAME:=db_clarine_hetic
EXEC:=docker exec

start:
	${DOCKER_COMPOSE} ${UP}

stop:
	${DOCKER_COMPOSE} ${DOWN}

prune:
	docker volume prune

db: start
	${EXEC} -ti ${CONTAINER_DB_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -w

insert-data: truncate
	${EXEC} -i ${CONTAINER_DB_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} < database/dump.sql

truncate:
	${EXEC} -i ${CONTAINER_DB_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"