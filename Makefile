SHELL:=/bin/bash
include .env
export $(shell sed 's/=.*//' .env)

# clean:
# 	rm -rf dist

run:
	go run main.go