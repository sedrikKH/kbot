VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

build: format
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${shell dpkg --print-architecture} go build -v -o prometheus_kbot -ldflags "-X="hgithub.com/sedrikKH/prometheus_kbot/cmd.appVersion=${VERSION}

clean:
	rm -rf prometheus_kbot 
