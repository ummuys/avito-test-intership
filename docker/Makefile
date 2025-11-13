update-image:
	-docker compose down 
	-docker image rm docker-pr-service
	-docker compose up -d


update-volumes:
	-docker compose down -v
	-docker compose up -d

update-all:
	-docker compose down -v
	-docker image rm docker-pr-service
	-docker compose up -d

check-containers:
	-docker compose ps -a
