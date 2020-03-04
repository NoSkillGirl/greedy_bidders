################################################################
#  BUILD
################################################################
FROM golang:1.14-alpine3.11 AS builder
ADD . /src
WORKDIR /src
RUN go mod download
RUN CGO_ENABLED=0 go build -o /service .


################################################################
#  MAIN
################################################################
FROM alpine:3.11

RUN echo "https://dl-4.alpinelinux.org/alpine/v3.11/main" >/etc/apk/repositories
RUN echo "https://dl-4.alpinelinux.org/alpine/v3.11/community" >>/etc/apk/repositories
RUN apk add --no-cache --update curl ca-certificates && update-ca-certificates

COPY --from=builder /service /service
COPY config.yaml /config.yaml
HEALTHCHECK CMD curl --fail http://0.0.0.0:8080/ || exit 1
ENTRYPOINT ["/service"]
EXPOSE 8080
