BINARY=bin/amsl

PHONY: build
build:
	set GOARCH=amd64&&set GOOS=linux&&set CGO_ENABLED=0&&go build -o ${BINARY}
