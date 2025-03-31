.PHONY: lab


# Variables
CEP=03031040

up:
	docker compose up -d;

down:
	docker compose down;

docker-clean-up:
	sudo docker rm -f $(docker ps -a -q)
	sudo docker rmi -f $(docker images -q)
	sudo docker system prune -af
	sudo docker volume prune -f

svc-a:
	@sleep 3s ;
	curl -X POST -d '{"cep": "$(CEP)"}' http://localhost:8080
	@echo '\n' ;

svc-b:	
	@sleep 3s ;
	curl http://localhost:8081/$(CEP)
	@sleep 3s ;
	@echo '\n' ;

lab: infra-up svc-a svc-b

run: svc-a svc-b
