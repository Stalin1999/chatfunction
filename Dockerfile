FROM golang:1.17

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

COPY go.mod /build
COPY go.sum /build/

RUN cd /build/ && git clone https://github.com/Stalin1999/chatfunction.git
RUN cd /build/chatfunction/server/ && go build ./...

EXPOSE 9080

ENTRYPOINT [ "/build/chatfunction/server/server" ]