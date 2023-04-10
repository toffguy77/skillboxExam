SRC_DIR = ./cmd
DEST_DIR = ./bin
NAME_SERVER = statusPage

.PHONY: all
all: clean test vendor build run

.PHONY: clean
clean:
	rm -f $(DEST_DIR)/statusPage-*

.PHONY: test
test:
	go test ./...

.PHONY: vendor
vendor: test
	go mod tidy
	go mod vendor

.PHONY: build
build: clean test vendor
	bin/go-executables-build.bash $(NAME_SERVER) $(SRC_DIR) $(DEST_DIR)

.PHONY: run
run: build
	bin/go-executables-run.bash
