.SHELL = /bin/zsh

.PHONY = usage build run clean 

all = usage

# Purely meant to help speed up debugging on a local machine

usage:
	@echo "Usage:"
	@echo "make build -> Compile go file to an executable"
	@echo "make run -> Builds executable and runs it"
	@echo "make clean -> Removes an unnecessary executables"

build: clean
	@go build main.go

run: build
	@./main

clean:
ifneq (,$(wildcard ./main))
	@rm ./main 
endif
ifneq (, $(wildcard ./a.out))
	@rm ./a.out
endif