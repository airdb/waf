SHELL = /bin/bash

VERSION:=$(shell git describe --dirty --always)
#VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse HEAD)
REPO := github.com/airdb/waf

LDFLAGS=-ldflags
LDFLAGS += "-X=github.com/airdb/sailor/version.Repo=$(REPO) \
            -X=github.com/airdb/sailor/version.Version=$(VERSION) \
            -X=github.com/airdb/sailor/version.Build=$(BUILD) \
            -X=github.com/airdb/sailor/version.BuildTime=$(shell date +%s)"

default: swag build deploy

build:
	GOOS=linux go build $(LDFLAGS) -o main main.go

swag:
	swag init --generalInfo main.go

dev: swag
	env=dev go run $(LDFLAGS) main.go
