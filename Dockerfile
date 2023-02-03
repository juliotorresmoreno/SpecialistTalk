FROM golang:1.19-alpine

ADD . /app

WORKDIR /app

RUN go mod download
RUN go build -o /SpecialistTalk

EXPOSE 4080

ENTRYPOINT [ "/SpecialistTalk" ]
