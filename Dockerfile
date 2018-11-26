FROM golang:latest

RUN apt update
RUN apt install -y \
    net-tools \
    netcat \
    telnet

EXPOSE 5233

ENTRYPOINT ["/go/src/scheduler/docker-files/tcp_server.sh", "5555"]

