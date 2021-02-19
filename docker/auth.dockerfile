FROM golang

COPY . /usr/code

WORKDIR /usr/code

RUN go build -o auth cmd/auth.go

ENTRYPOINT ./auth