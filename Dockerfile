# compile
FROM golang:alpine as builder

WORKDIR /data
ENV GOPROXY=https://goproxy.cn,direct

COPY ./go.mod ./go.mod
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add git && \
    go mod download

COPY . .
RUN mkdir build && \
    cp -R web/views ./build && \
    go build -o ./build/app main.go
    
# build image
FROM alpine

MAINTAINER zhangsunbaohong <zhangsunbaohong@163.com>
LABEL version="0.1.0"

WORKDIR /data

COPY --from=builder /data/build .

EXPOSE 8080
ENTRYPOINT [ "./app" ]

