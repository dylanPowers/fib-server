FROM ubuntu:14.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      ca-certificates \
      curl

ENV GO_VERSION 1.6
RUN curl https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz | \
      tar -vxz -C /opt && \
    ln -s /opt/go/bin/go /usr/local/bin/
ENV GOROOT /opt/go

RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY fib /opt/fib

RUN useradd --system --shell /usr/sbin/nologin fib
USER fib

EXPOSE 8080
CMD go run main.go
