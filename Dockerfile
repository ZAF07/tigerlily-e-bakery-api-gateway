FROM golang:1.17 as dev

WORKDIR /opt/tigercoders/bff

RUN apt update && apt upgrade -y && \
  apt install -y git \
  make openssh-client

RUN git config --global url."git@github.com".insteadof "https://github.com/"

# RUN go install github.com/cosmtrek/air@latest

# Have to use a older version of air since i am using GO1.17(they require GO1.22)
RUN go install github.com/cosmtrek/air@b3a0f1348a7584b7be0dbe9a7a61d63d6f181ce8

CMD ["air"]