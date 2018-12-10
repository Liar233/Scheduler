FROM golang:latest

RUN apt update
RUN apt install -y \
    net-tools \
    netcat \
    telnet

EXPOSE 5233

ENTRYPOINT ["go", "run", "./docker-files/tcp_server.go", "5555"]

