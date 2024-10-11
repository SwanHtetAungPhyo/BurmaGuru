SRC=./go-api/main.go
TARGET=./bin/api

.PHONY: build
build:
	@echo "Building the application ..."
	@echo "API in listening at the http://localhost:8080/api/"
	go build -o $(TARGET) $(SRC)

.PHONY: run
run:
	@echo "Run the application ..."
	go run $(SRC)


.PHONY: clean
clean:
	@echo "clean the application..."
	rm -rf $(TARGET)

.PHONY: deps
deps:
	@echo "Installing the depedencies..."
	go mod tidy
	go mod vendor

.PHONY: docs
docs:
	aglio -i api.apib -o api-docs.html
