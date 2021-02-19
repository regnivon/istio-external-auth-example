FROM golang

COPY . /usr/code

WORKDIR /usr/code

RUN go build -o main main.go

ENTRYPOINT ./main