FIB_PORT = 8080

run: build
	if ! docker ps -a | grep -w fib-server -q; then  \
		docker create                                  \
					--name fib-server                        \
					--publish $(FIB_PORT):8080               \
					--restart unless-stopped                 \
					-v /etc/timezone:/etc/timezone           \
					--workdir /opt/fib/                      \
					fib-server;                              \
	else																						 \
		$(MAKE) stop;                                  \
	fi
	docker cp src fib-server:/opt/fib/src
	docker start fib-server


build:
	if ! docker images | grep -w fib-server -q; then  \
		$(MAKE) force-build;                            \
	fi


force-build: remove-all
	docker build --tag fib-server .


remove-all: stop
	docker rm fib-server 2> /dev/null | true
	docker rmi fib-server 2> /dev/null | true


stop:
	docker stop fib-server | true


test: build
	docker run                          \
		--rm                              \
		--workdir /opt/fib                \
		-v $(shell pwd)/src:/opt/fib/src  \
		fib-server                        \
		go test .


.PHONY: run build force-build remove-all stop test
