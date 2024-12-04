up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

run-main:
	go run main.go & go run consumer/consumer.go

	
