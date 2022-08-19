FROM golang:1.17 as dev

WORKDIR /opt/tigercoders/bff

RUN apt update && apt upgrade -y && \
  apt install -y git \
  make openssh-client

RUN git config --global url."git@github.com".insteadof "https://github.com/"

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]