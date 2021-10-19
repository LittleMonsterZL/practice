export tag=v1.0
root:
	export ROOT=github.com/LittleMonsterZL/practice

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t practice/httpserver:${tag} .

push: release
	echo "pushing practice/httpserver"
	docker push practice/httpserver:${tag}

pull:
	sudo docker pull practice/httpserver:${tag}
run:
	sudo docker run -itd practic/httpserver:${tag}
