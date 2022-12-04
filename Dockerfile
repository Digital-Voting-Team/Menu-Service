FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/menu-service
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/menu-service /go/src/menu-service


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/menu-service /usr/local/bin/menu-service
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["menu-service"]
