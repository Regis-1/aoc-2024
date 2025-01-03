# Makefile for Go side of Advent of Code 2024
all: binaries

binaries: day2 day4

.PHONY: day2
day2:
	@go build -o bin/day2 cmd/day2/main.go

.PHONY: day4
day4:
	@go build -o bin/day4 cmd/day4/main.go

.PHONY: clean
clean:
	@rm -rf bin/*
