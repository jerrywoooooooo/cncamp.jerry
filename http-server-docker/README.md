构建本地镜像,编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
```shell
docker build --no-cache -f http-server.Dockerfile -t jerrywoooooooo/http-server:v1.0 ./
```
将镜像推送至 Docker 官方镜像仓库。
```shell
# Docker ID : jerrywoooooooo
docker login 
docker push  jerrywoooooooo/http-server:v1.0
```
通过 Docker 命令本地启动 httpserver。
```shell
docker run -p 8080:80 -d jerrywoooooooo/http-server:v1.0
# 验证
curl http://127.0.0.1:8080/healthz
```
通过 nsenter 进入容器查看 IP 配置。
```shell
# 查看ip
docker inspect --format "{{ .State.Pid }}" `docker ps | grep  jerrywoooooooo/http-server:v1.0 | awk '{print $1}'` | xargs  -i -t nsenter -t {} -n ip addr show 

# or 

# get container id
docker ps | grep  jerrywoooooooo/http-server:v1.0 | awk '{print $1}'
# get pid
docker inspect --format "{{ .State.Pid }}" <containerid>
# show ip
nsenter -n ip addr show -t <pid> 
```