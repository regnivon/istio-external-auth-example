FROM golang

COPY . /usr/code

WORKDIR /usr/code

RUN go build -o server cmd/server.go

ENTRYPOINT ./server