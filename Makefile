.PHONY: default help
	
default: help

help:
	echo "test"

build:
	go build -o bin/main main.go

run:
	./bin/main
