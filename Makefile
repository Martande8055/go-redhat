APP_NAME = average-calculation-api
BINARY = main
PORT = 3100


GO_BUILD = go build -o $(BINARY) main.go
GO_TEST = go test ./... -v
GO_CLEAN = go clean

.PHONY: all build run test clean docker-build docker-run

all: 
	make clean
	make test
	make build
	make run

build:
	$(GO_BUILD)

run: build
	./$(BINARY)

test:
	$(GO_TEST)

clean:
	$(GO_CLEAN)
	rm -f $(BINARY)

podman-build:
	podman build -t $(APP_NAME) .

podman-run:
	podman run -p $(PORT):$(PORT) $(APP_NAME)

changelog:
	@changelog.sh