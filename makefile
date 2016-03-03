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
	docker cp src/fib fib-server:/opt/fib
	docker start fib-server


build:
	if ! docker images | grep -w fib-server -q; then  \
		$(MAKE) force-build;                            \
	fi


force-build:
	docker build --tag fib-server src/


remove-all: stop
	docker rm fib-server
	docker rmi fib-server


stop:
	docker stop fib-server


test: build
	docker run                          \
		--rm                              \
		--workdir /opt/fib                \
		-v $(shell pwd)/src/fib:/opt/fib  \
		fib-server                        \
		go test .


.PHONY: run build force-build remove-all stop test
