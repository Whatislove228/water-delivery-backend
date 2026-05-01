APP_NAME=water-delivery
DB_URL=postgres://postgres:postgres@localhost:5432/water_delivery?sslmode=disable

run:
	go run ./cmd/api

test:
	go test ./...

docker-up:
	docker compose up --build

docker-down:
	docker compose down

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down 1

seed-products:
	docker compose exec -T postgres psql -U postgres -d water_delivery < ./seeds/products.sql