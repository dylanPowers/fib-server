run_image = fib
test_image = fib_test


run: build
	if ! docker ps -a | grep -w fib-server -q; then \
	  docker create                                 \
	        --name fib-server                       \
	        --restart unless-stopped                \
	        -v /etc/timezone:/etc/timezone          \
	        --workdir /opt/fib/                     \
	        fib-server;                             \
	else																						\
		$(MAKE) stop                                  \
	fi
	docker cp src/fib fib-server:/opt/fib
	docker start fib-server


build:
	if ! docker images | grep -w fib-server -q; then \
		$(MAKE) force-build;                           \
	fi


force-build:
	docker build --tag fib-server src/


stop:
	docker stop fib-server



.PHONY: run build force-build stop
