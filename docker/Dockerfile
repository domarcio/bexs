FROM golang:1.15

RUN mkdir -p /app
WORKDIR /app
COPY . /app

EXPOSE 7007

RUN go mod download