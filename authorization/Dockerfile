FROM golang:alpine3.13 as builder
RUN mkdir /go/src/authorization
RUN apk add git
COPY . /go/src/authorization
WORKDIR /go/src/authorization
RUN go mod tidy
RUN go build -o /go/bin/authorization .



FROM alpine:latest
RUN mkdir /authorization
WORKDIR /authorization
COPY --from=builder /go/bin/authorization .
RUN chmod 755 authorization
CMD ["./authorization"]
