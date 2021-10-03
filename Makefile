up: build-docker-migrate build-docker-server
	docker-compose up -d

down:
	docker-compose down

migrate:
	goose -dir migrations/sql mysql "pay_sys:pay_sys_passw@tcp(0.0.0.0:3306)/payment_system?parseTime=true" up
.PHONY: migrate

build-docker-migrate:
	docker build -t ps-test/migration ./migrations
.PHONY: build-docker-migrate

build-docker-server:
	docker build -t ps-test/server -f server/Dockerfile ./server

run-server:
	source ./configs/.env.local && cd server && go run . server
.PHONY: run-server

report:
	source ./configs/.env.local && cd server && go run . report --user_id=$(user_id) --begin_time=$(begin_time) --end_time=$(end_time) --filename=$(filename)
.PHONY: report
