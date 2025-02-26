main:
	mkdir build
	go build -o build/uniklub-server

clean:
	rm -r build

migrate-up:
	go run -tags='pgx5' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path ./database/migrations/ -database pgx5://postgres:postgres@127.0.0.1:5432/app up

migrate-down:
	go run -tags='pgx5' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path ./database/migrations/ -database pgx5://postgres:postgres@127.0.0.1:5432/app down

migrate: migrate-up

all: main
