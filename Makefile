# Copyright 2017 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 2>/dev/null)
BUILD_DATE = $(shell date -u +%Y%m%d-%H:%M:%S)

LDFLAGS += -X transwarp/go-web-installer/pkg/version.Version=${GIT_TAG}
LDFLAGS += -X transwarp/go-web-installer/pkg/version.GitSha1Version=${GIT_SHA}
LDFLAGS += -X transwarp/go-web-installer/pkg/version.BuildDate=${BUILD_DATE}

#.PHONY: go-swagger
#go-swagger:
#	swagger generate spec -o ./swagger.json

.PHONY: build
build:
	go build -ldflags '$(LDFLAGS)'

all: build
