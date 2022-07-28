SHELL := /bin/bash

# ======================================================================================================================
# Call solutions

atoi:
	go run main.go atoi

regexpMatch:
	go run main.go regexpMatch

help:
	go run main.go

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
