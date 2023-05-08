init:
	docker-compose up --build -d

stop:
	docker-compose down

start:
	docker-compose up

reset:
	sudo rm -r .database/
