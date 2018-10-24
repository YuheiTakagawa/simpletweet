CMD=	go
GOBUILD=$(CMD) build
GOCLEAN=$(CMD) clean
GOTEST=	$(CMD) test
GOGET=	$(CMD) get
NAME=	simpletweet


BIN=	/usr/local/bin

.PHONY: all
all: build
	cp $(NAME) $(BIN)

.PHONY: deps
build: deps
	$(GOBUILD) -o $(NAME) -v

.PHONY: test
test: deps
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(NAME)

.PHONY: run
run:
	$(GOBUILD) -o $(NAME) -v ./...
	./$(NAME)

.PHONY: deps
deps:
	$(GOGET) github.com/garyburd/go-oauth/oauth
	$(GOGET) github.com/jessevdk/go-flags
	
