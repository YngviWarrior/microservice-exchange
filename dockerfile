FROM golang:1.24rc3-alpine

WORKDIR /app

COPY . /app/

RUN go mod tidy

EXPOSE 3001
ENTRYPOINT [ "go","run","main.go" ]