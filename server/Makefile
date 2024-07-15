BINARY_NAME=udp-server
BUILD_DIR=bin
SRC_DIR=.
ARCH=arm64
VERSION=$(shell git describe --tags --always)
COMMIT=$(shell git rev-parse HEAD)
DATE=$(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS="-X main.Version=${VERSION} -X main.Commit=${COMMIT} -X main.Date=${DATE}"

all: build

build:
	@echo "Building the binary..."
	GOOS=linux GOARCH=${ARCH} go build -ldflags ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME} ${SRC_DIR}

clean:
	@echo "Cleaning the build directory..."
	rm -rf ${BUILD_DIR}

test:
	@echo "Running tests..."
	go test ./...

deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download

install: build
	@echo "Installing the binary..."
	go install ${SRC_DIR}

run:
	@echo "Running the application..."
	go run ${SRC_DIR}

.PHONY: all build clean test deps install run
