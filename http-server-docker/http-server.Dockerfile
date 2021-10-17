# 编译环境
FROM golang:latest as builder

LABEL version="v1.0" description="极客时间云原生训练营Docker小练习" by="jerrywoooooooo" source="https://github.com/jerrywoooooooo/cncamp.jerry.git"

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /home/app/go

COPY build.sh ./
RUN chmod +x build.sh && sh build.sh

# 运行环境
FROM ubuntu:latest as runner

WORKDIR /home/app/

COPY --from=builder /home/app/go/cncamp.jerry/http-server/http-server .
EXPOSE 80
ENTRYPOINT ["./http-server"]


