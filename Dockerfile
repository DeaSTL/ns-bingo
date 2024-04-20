FROM golang:alpine

RUN apk add make

RUN mkdir go

ENV GOPATH=/go

RUN go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

COPY . .

RUN make build


ENTRYPOINT make run
