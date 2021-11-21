export tag=v1.1
root:
	export ROOT=github.com/cncamp.jerry/http-server

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release:
	echo "building httpserver container"
	docker build --no-cache -t jerrywoooooooo/http-server:${tag} ./

push: release
	echo "pushing cncamp/httpserver"
	docker push  jerrywoooooooo/http-server:${tag}
