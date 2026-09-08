NAME := ./bin/frozen
SRC := ./cmd/frozen
TESTS := ./tests/...

PORT := 7000

run: build
	@$(NAME) $(PORT)

build:
	@go build -o $(NAME) $(SRC)

test:
	go test $(TESTS)

clean:
	rm -rf $(NAME)