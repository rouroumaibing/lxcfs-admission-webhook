FROM golang:1.17-alpine3.15 as build

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /src

ADD . .

RUN GOOS=linux go build -o lxcfs-admission-webhook ./cmd/

FROM alpine:3.15

WORKDIR /lxcfs

COPY --from=build /src/lxcfs-admission-webhook /lxcfs/lxcfs-admission-webhook

ENTRYPOINT ["/lxcfs/lxcfs-admission-webhook"]
