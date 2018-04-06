.PHONY: world compress otb-bus

otb-bus:
	CC=/usr/sbin/musl-gcc CGO_ENABLED=0 GOOS=linux go build -ldflags="-linkmode external -extldflags "-static"" -v -o otb-bus bin/otb-bus/*.go

compress:
	find -P . -maxdepth 1 -name 'otb-*' -type f -executable -exec upx --ultra-brute {} \;

world: otb-bus

all: world
