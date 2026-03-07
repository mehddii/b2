BIN_NAME = b2
MAIN_PACKAGE_PATH = .

.PHONY: run
run: build
	@/tmp/bin/$(BIN_NAME)

.PHONY: build
build: clean tidy fmt vet
	@go build -v -o /tmp/bin/$(BIN_NAME) $(MAIN_PACKAGE_PATH)

.PHONY: test
test:
	@go test -race $(MAIN_PACKAGE_PATH)

.PHONY: clean
clean:
	@rm -rf /tmp/bin/$(BIN_NAME) \
		&& go clean ./...

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: vet
vet:
	@go vet ./...
