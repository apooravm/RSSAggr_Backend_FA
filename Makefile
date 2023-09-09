.PHONY: build run clean tidy

APP_NAME := RSSAggr_Backend_FA.exe
rand := bruh

build:
	go build -o bin/$(APP_NAME)

run: build
	./bin/$(APP_NAME)

clean:
	rm -f $(APP_NAME)

tidy:
	go mod tidy && go mod vendor

init:
	go mod init github.com/apooravm/RSSAggr_Backend_FA && go get
