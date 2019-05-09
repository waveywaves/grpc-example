
#base
FROM golang:1.11-alpine AS base
LABEL maintainer="Vibhav Bobade <vibhav.bobde@gmail.com>"

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR $GOPATH/src/github.com/waveywaves/grpc-example/

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod .
COPY go.sum .

RUN go mod download

# #builder
# FROM golang:1.11-alpine AS builder
# COPY --from=base . .
# WORKDIR $GOPATH/src/github.com/waveywaves/grpc-example/
COPY . .
RUN go install -a -tags netgo -ldflags '-w -extldflags "-static"' github.com/waveywaves/grpc-example

#main
FROM alpine AS grpc-example
COPY --from=base /go/bin/grpc-example /bin/grpc-example

EXPOSE 3000
ENTRYPOINT /bin/grpc-example