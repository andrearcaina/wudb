.PHONY: default help
	
default: help

help:
	echo "test"

build:
	go build -o bin/main .

run:
	./bin/main $(ARGS)
