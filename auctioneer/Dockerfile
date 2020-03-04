################################################################
#  BUILD
################################################################
FROM golang:1.14-alpine3.11 AS builder
COPY go.mod /src/go.mod
COPY go.sum /src/go.sum
WORKDIR /src
RUN go mod download
ADD . /src
RUN CGO_ENABLED=0 go build -o /service .


################################################################
#  MAIN
################################################################
FROM alpine:3.11
RUN echo "https://dl-4.alpinelinux.org/alpine/v3.11/main" >/etc/apk/repositories
RUN echo "https://dl-4.alpinelinux.org/alpine/v3.11/community" >>/etc/apk/repositories
RUN apk add --no-cache --update curl ca-certificates && update-ca-certificates
RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

COPY --from=builder /service /service
COPY config.yaml /config.yaml
HEALTHCHECK CMD curl --fail http://0.0.0.0:8080/ || exit 1
# ENTRYPOINT ["/service"]
EXPOSE 8080
