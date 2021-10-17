### jerrywoo的极客时间云原生训练营课后练习
### G20210797010117 / jerrywoooooooo

作业一：
#### http-server
1. 接收客户端request，并将request中带的header写入response header
2. 读取当前系统的环境变量中的VERSION配置，并写入response header
3. Server端记录访问日志包括客户端ip, HTTP返回码, 输出到server端的标准输出
4. 当访问localhost/healthz时, 应返回200

作业二：
#### http-server & docker
1. 构建本地镜像。
2. 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
3. 将镜像推送至 Docker 官方镜像仓库。
4. 通过 Docker 命令本地启动 httpserver。
5. 通过 nsenter 进入容器查看 IP 配置。
