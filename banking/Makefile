include .env
export

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help


swag-v1: ### swag init
	swag.exe init -g app/app.go
.PHONY: swag-v1

run: swag-v1 ### swag run
	go mod tidy && \
	 go run main.go
.PHONY: run