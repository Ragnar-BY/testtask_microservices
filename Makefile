.PHONY: build clean lint lintd run

build: build_web build_players

build_web:
	go build -o ./bin/WebServer github.com/Ragnar-BY/testtask_microservices/webserver

build_players:
	go build -o ./bin/playerServiceBin github.com/Ragnar-BY/testtask_microservices/playerService

clean:
	rm -rf ./bin

lint:
	gometalinter --enable=misspell --enable=unparam --enable=dupl --enable=gofmt --enable=goimports --disable=gotype --disable=gas --deadline=3m ./...

lintd:
	gometalinter --enable=misspell --enable=unparam --enable=gofmt --enable=goimports --disable=gotype --disable=gas --deadline=3m ./...

mongo:
	docker start mongodb

run: mongo
	./bin/playerServiceBin & ./bin/WebServer

