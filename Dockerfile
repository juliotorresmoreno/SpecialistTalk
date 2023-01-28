FROM golang:1.19-alpine

ADD . /app

WORKDIR /app

RUN go mod download
RUN go build -o /freelive

EXPOSE 4080

ENTRYPOINT [ "/freelive" ]
