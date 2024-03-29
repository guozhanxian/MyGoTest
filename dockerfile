#源镜像
FROM golang:latest
#作者
MAINTAINER Ralph "guozhanxian@gmail.com"
#设置工作目录
WORKDIR $GOPATH/src/MyGoTest
#设置环境变量
ENV GO112MODULE ON
ENV GOPROXY https://goproxy.io

#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/MyGoTest
#go构建可执行文件
RUN go build ./main.go
#暴露端口
EXPOSE 8199
#最终运行docker的命令
ENTRYPOINT  ["./main"]
