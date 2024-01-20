APP := $(shell basename $(shell git remote get-url origin))
REGISTRY=ghcr.io/sedrikkh
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux
TARGETARCH=amd64

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

get:
	go get

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags "-X="github.com/sedrikKH/kbot/cmd.appVersion=${VERSION}


linux:
	TARGETOS=linux
	TARGETARCH=arm64
	make build

arm:
	TARGETOS=linux
	TARGETARCH=arm
	make build

macos:
	TARGETOS=darwin
	TARGETARCH=amd64
	make build

windows:
	TARGETOS=windows
	TARGETARCH=amd64
	make build


image:
# --build-arg BOT_TOKEN=${TELE_TOKEN}
	docker build  . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

clean:
	rm -rf kbot 
	docker rmi ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}
