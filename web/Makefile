BINARY_NAME=udp-client
BUILD_DIR=bin
SRC_DIR=.
ARCH=arm64
LOCAL_OS=darwin

all: build

build:
	@echo "Building the binary..."
	GOOS=${LOCAL_OS} GOARCH=${ARCH} go build -o ${BUILD_DIR}/${BINARY_NAME} ${SRC_DIR}

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

