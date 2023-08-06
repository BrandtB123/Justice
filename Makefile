# Build Docker images
build:
	docker build -t my-justice-app .   

# Start containers
up:
	docker compose -f build/package/docker-compose.yml up -d

# Stop containers
down:
	docker compose -f build/package/docker-compose.yml down