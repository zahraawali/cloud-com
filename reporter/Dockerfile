FROM golang:1.18-alpine

WORKDIR /src/app
COPY . .

RUN go mod download

EXPOSE 8081

CMD ["go", "run", "main.go", "serve"]
