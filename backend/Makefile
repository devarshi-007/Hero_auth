.PHONY: create-migration start migrateup migratedown
create-migration:
ifneq (, $(@shell ./migrate -version))
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz
	mv migrate.linux-amd64 migrate
	chmod +x migrate
endif
	migrate create -ext sql -dir database/migrations $(file_name)

# build:
# 	go build -o=$(app_name) .

start:
	go run app.go .

migrateup:
	migrate -path database/migration -database "postgres://devarshi:eA234123Pq@localhost:5432/supreme_org?sslmode=disable" -verbose up

# start-seed:
# 	go run app.go load-data

migratedown:
	migrate -path database/migration -database "postgres://devarshi:eA234123Pq@localhost:5432/supreme_org?sslmode=disable" -verbose down
# start-dev:
# 	@nodemon --exec go run app.go api --signal SIGTERM