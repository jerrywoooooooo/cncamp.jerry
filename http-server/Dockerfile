# 编译环境
FROM golang:alpine3.14 as builder

LABEL version="v1.1" description="极客时间云原生训练营Docker小练习" by="jerrywoooooooo" source="https://github.com/jerrywoooooooo/cncamp.jerry.git"

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /home/app/httpserver

COPY . .
RUN go build -mod=mod .

# 运行环境
FROM alpine:3.14 as runner

WORKDIR /home/app/

COPY --from=builder /home/app/httpserver/ .
EXPOSE 80
ENTRYPOINT ["./httpserver"]


