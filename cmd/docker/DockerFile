FROM golang:1.22-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download
RUN go mod verify

COPY . .

RUN go build -o main main.go

EXPOSE 3000

CMD ["./main"]
