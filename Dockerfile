FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/Menu-Service
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/Menu-Service /go/src/Menu-Service


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/Menu-Service /usr/local/bin/Menu-Service
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["Menu-Service"]
