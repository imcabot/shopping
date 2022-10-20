FROM golang:alpine AS builder

#设置必要的环境变量
ENV GO111MODULE=on\
    GGOS=linux\
    GOARCH="amd64"\
    GOProxy="https://goproxy.cn,direct"\
    CGO_ENABLED=0

#移动到工作目录
WORKDIR /app

# 复制项目中的go.mod和go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

#将代码编译成二进制可执行文件
RUN go build -o shopping .

#创建一个小镜像
FROM scratch

#静态文件
COPY ./config /config


# 从builder镜像中把shopping拷贝到当前目录
COPY --from=builder /app/shopping /

EXPOSE 8080

# 需要运行的命令
ENTRYPOINT ["/shopping"]
