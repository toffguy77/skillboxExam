SRC_DIR = ./cmd
NAME_SERVER = statusPage




.PHONY: all
all: clean test vendor build

.PHONY: clean
clean:
	if [ -f $(NAME_SERVER) ]; then rm $(NAME_SERVER); fi

.PHONY: test
test:
	go test ./...

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: build
build:
	go build -o $(NAME_SERVER) $(SRC_DIR)
