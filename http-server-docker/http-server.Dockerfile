FROM golang:latest

LABEL version="v1.0" description="极客时间云原生训练营Docker小练习" by="jerrywoooooooo" source="https://github.com/jerrywoooooooo/cncamp.jerry.git"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /home/app/http-server

COPY run.sh run.sh
RUN chmod +x run.sh

CMD ["./run.sh"]