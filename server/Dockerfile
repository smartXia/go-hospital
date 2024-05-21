FROM  golang:1.22.3-alpine3.19 AS builder
# 作者
MAINTAINER xiapeifu
# 全局工作目录
WORKDIR /root
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY ./server /root
# 环境变量
#  用于代理下载go项目依赖的包
ENV GOPROXY https://goproxy.cn,direct
ENV GOCACHE /data/go/cache
RUN go mod download
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o devops



# 使用alpine这个轻量级镜像为基础镜像--运行阶段
# 使用 alpine 镜像作为最终基础镜像，因为它轻量级且易于设置时区
FROM alpine:latest
# 全局工作目录
WORKDIR /root
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /root/devops .

COPY --from=builder /root/config.yaml /root/config.yaml
COPY --from=builder /root/entrypoint.sh /root/entrypoint.sh

# 设置时区环境变量
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# 安装 tzdata 包以支持自定义时区设置
RUN apk --no-cache add tzdata
#设置时区
ENV TZ=Asia/Shanghai

RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime
# 暴露端口
EXPOSE 8334

# 可外挂的目录
#CMD ["entrypoint.sh"]
RUN chmod +x /root/entrypoint.sh
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["/bin/sh", "/root/entrypoint.sh"]