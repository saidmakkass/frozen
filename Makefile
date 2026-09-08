NAME := ./bin/frozen
SRC := ./cmd/frozen
TESTS := ./tests/...

PORT := 7000

run: $(NAME)
	@$(NAME) $(PORT)

build: $(NAME)

$(NAME):
	@go build -o $(NAME) $(SRC)

test:
	go test $(TESTS)

clean:
	rm -rf $(NAME)

.PHONY: run build test clean