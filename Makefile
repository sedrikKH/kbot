APP=$(shell basename $(shell git remote get-url origin))
REGISTRY=sedrikkh
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
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o prometheus_kbot -ldflags "-X="hgithub.com/sedrikKH/prometheus_kbot/cmd.appVersion=${VERSION}


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
	docker build --build-arg BOT_TOKEN=${TELE_TOKEN} . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

clean:
	rm -rf prometheus_kbot 
	docker rmi ${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}
