# 引入go镜像作为编译环境
FROM golang as builder

# 设置变量app_env
ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/opengateway
WORKDIR /go/src/opengateway

# 编译
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o opengateway .

# 再次引入go镜像用作运行环境，也可以换成其他的镜像，这里使用alpine镜像
FROM alpine
WORKDIR /app
COPY --from=builder /go/src/opengateway/opengateway .
COPY --from=builder /go/src/opengateway/serverConfig.yaml .

# 运行服务
CMD mkdir /lib64
CMD ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
CMD ./opengateway

# 对外开放端口
EXPOSE 8000
