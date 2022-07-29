SHELL := /bin/bash

# ======================================================================================================================
# Commands for benchmarking
# go test -run none -bench . -benchmem -memprofile p.out
# go test -run none -bench . -benchmem -cpuprofile p.out
# go tool pprof containerWithMostWater.test p.out

# ======================================================================================================================
# Call solutions

atoi:
	go run main.go atoi

regexpMatch:
	go run main.go regexpMatch

isPalindromeNum:
	go run main.go isPalindromeNum

reverseInt:
	go run main.go reverseInt

cwmw:
	go run main.go cwmw

help:
	go run main.go

# ======================================================================================================================
# Running tests within the local computer

test:
	go test ./... -count=1

test1:
	go test -v ./...

check:
	staticcheck -checks=all ./...

# ======================================================================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all
