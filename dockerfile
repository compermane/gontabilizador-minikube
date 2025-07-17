FROM golang:1.23.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd
RUN chmod +x main

EXPOSE 8000

CMD ["./main"]
