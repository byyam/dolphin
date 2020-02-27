GOOS=darwin
ifeq ($(os),linux)
	GOOS=linux
endif


all: http-server tcp-server

lite: http-server tcp-server

http-server:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -mod=vendor -o bin/http-server ./service/http-server

tcp-server:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -mod=vendor -o bin/tcp-server ./service/tcp-server


pb:
	protoc ./proto/hello.proto --go_out=plugins=grpc:./

clean:
	rm -f ./bin/*
