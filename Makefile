#!make
include .env

binary_path = dist/upprove-mcp-server
main_path = cmd/main.go

# Server
build:
	go build -o $(binary_path) $(main_path)

run-dev:
	go run $(main_path)

run:
	$(binary_path)

fmt-check:
	gofmt -l .

# Database
db-start:
	docker compose up -d

db-connect-admin:
	docker exec -it upprove-db mongosh admin -u ${MONGO_INITDB_ROOT_USERNAME} -p ${MONGO_INITDB_ROOT_PASSWORD}

db-create-user:
	docker exec -it upprove-db \
	sh -c 'mongosh admin -u ${MONGO_INITDB_ROOT_USERNAME} -p ${MONGO_INITDB_ROOT_PASSWORD} \
	--eval "use $(UPPROVE_DB)" \
	--eval "db.createUser({user: \"$(UPPROVE_USER)\", pwd: \"$(UPPROVE_PWD)\", roles: [{role: \"readWrite\", db: \"$(UPPROVE_DB)\"}]})"'

db-connect-user:
	docker exec -it upprove-db mongosh "mongodb://${UPPROVE_USER}:${UPPROVE_PWD}@localhost:27017/upprove?authSource=${UPPROVE_DB}"