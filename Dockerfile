# build project using golang
FROM golang:1.19-alpine AS builder
ENV GOPROXY https://goproxy.cn/
COPY . /build_dir
WORKDIR /build_dir
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o hellosvr .

# package docker using alpine
FROM alpine
RUN apk add -U tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
COPY --from=builder /build_dir/hellosvr /
EXPOSE 8080

CMD ["/hellosvr", "server"]