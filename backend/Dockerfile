FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go generate ./...
RUN go build -o laas ./cmd/laas

EXPOSE 8080

CMD ["./laas"]
