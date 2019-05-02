
#builder
FROM golang:1.10 AS builder
LABEL maintainer="Vibhav Bobade <vibhav.bobde@gmail.com>"
RUN go get -v github.com/golang/dep/cmd/dep
WORKDIR $GOPATH/src/github.com/waveywaves/grpc-example/
ADD Gopkg.* ./
RUN dep ensure --vendor-only
ADD . .
RUN go install github.com/waveywaves/grpc-example

#obj-cache
FROM golang:1.10 as obj-cache
COPY --from=builder /root/.cache /root/.cache

#main
COPY --from=builder /go/bin/grpc-example /go/bin/grpc-example
EXPOSE 3000
ENTRYPOINT /go/bin/grpc-example