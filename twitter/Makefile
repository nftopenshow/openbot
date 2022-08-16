# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build -buildvcs=false
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

all: clean mac linux

mac:
	$(GOBUILD) -o ./bin/twitter-osx -v ./src/main/

test:
	$(GOTEST) -v ./...

clean:
	rm -f ./bin/twitter-*
	rm -fr ./bin/logs/*.log

clean-logs:

linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD)  -o ./bin/twitter-linux -v ./src/main/
