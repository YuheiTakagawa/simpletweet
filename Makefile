.PHONY: all

BIN=	/usr/local/bin
all:
	go build
	cp simpletweet ${BIN}
