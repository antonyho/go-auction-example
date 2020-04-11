GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
TARGET=auction-api-example

all: build

.PHONY: api test benchmark run image build

api:
	docker run --rm \
	-v $(PWD):/local openapitools/openapi-generator-cli generate \
	-i /local/resources/api/spec/v1/swagger.json \
	-g go \
	-o /local/openapi

test:
	$(GOTEST) -v -race -cover ./...

benchmark:
	$(GOTEST) -race -bench=. ./...

build:
	$(GOBUILD) -o $(TARGET) ./cmd

clean:
	rm $(TARGET)
	$(GOCLEAN)

image:
	docker build --rm -t antonyho-auction-api-example .

run:
	$(GORUN) cmd/main.go
