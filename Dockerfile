FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/sago35/grpczip
COPY . /usr/local/go/src/github.com/sago35/grpczip
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/grpczip ./grpczip


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/sago35/grpczip/build/grpczip /bin/grpczip
CMD ["grpczip", "up"]
