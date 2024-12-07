up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

run:
	go run cmd/app/main.go
	
