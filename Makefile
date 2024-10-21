APP_NAME = average-calculation-api
BINARY = main
PORT = 3100

GO = go
GO_BUILD = $(GO) build -o $(BINARY) main.go
GO_TEST = $(GO) test ./...
GO_CLEAN = $(GO) clean

all: build

build:
	$(GO_BUILD)

run: build
	./$(BINARY)

test:
	$(GO_TEST)

clean:
	$(GO_CLEAN)
	rm -f $(BINARY)

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run -p $(PORT):$(PORT) $(APP_NAME)

.PHONY: all build run test clean docker-build docker-run
