default:
	@echo "Building API..."
	docker build -f cmd/api/Dockerfile -t api .

up: default
	@echo "Starting API..."
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

clean: down
	@echo "Cleaning up..."
	rm -f api
	docker system prune -f
	docker volume prune -f