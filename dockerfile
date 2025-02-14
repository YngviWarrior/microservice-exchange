FROM golang:1.24rc3-alpine

WORKDIR /app

COPY . /app/

RUN go mod tidy
RUN go install github.com/air-verse/air@latest

EXPOSE 3001
ENTRYPOINT [ "go", "run", "." ]