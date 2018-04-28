export GOPATH=$(shell pwd)

all: livingserver 

livingserver:
	go build -o ./bin/livingserver ./src/livingserver/main.go
clean:
	rm -rf ./bin/*	\
