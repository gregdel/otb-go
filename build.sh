#!/bin/sh
# shellcheck disable=SC2086

export CGO_ENABLED=0
export CC=/usr/sbin/musl-gcc
export GOOS=linux

# Static build
# go build -ldflags="-linkmode external -extldflags "-static"" -v
# Compression
# find -P . -maxdepth 1 -name 'otb-*' -type f -executable -exec upx --ultra-brute {} \;

_build() {
	name=$1
	echo "Building $name..."
	go generate bin/$name/*.go
	go build -v -o "$name" bin/$name/*.go
	echo "Done bulding $name!"
}

_build otb-web-ui
