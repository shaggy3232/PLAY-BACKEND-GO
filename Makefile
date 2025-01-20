# fmt all golang code
.PHONY: fmt 
fmt:
	go fmt ./...

# run all golang tests
.PHONY: test
test:
	go test -v ./...

# run golang application locally
.PHONY: run
run:
	go run ./cmd/main/main.go

# build images and run all services in docker
.PHONY: compose
compose:
	docker compose up --build

# stop running docker containers
.PHONY: clean
clean:
	docker compose down
