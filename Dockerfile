ARG VERSION=1.15

FROM golang:${VERSION} as builder

WORKDIR /go/src/app
COPY . /go/src/app

RUN go build -tags netgo -o hello .

FROM alpine
COPY --from=builder /go/src/app/hello /usr/bin/hello
ENTRYPOINT ["/usr/bin/hello"]
EXPOSE 8080
