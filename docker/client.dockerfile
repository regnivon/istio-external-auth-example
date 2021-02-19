FROM golang

COPY . /usr/code

WORKDIR /usr/code

RUN go build -o client cmd/client.go

ENTRYPOINT ./client