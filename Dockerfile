FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/nonce-auth-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/nonce-auth-svc /go/src/nonce-auth-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/nonce-auth-svc /usr/local/bin/nonce-auth-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["nonce-auth-svc"]
