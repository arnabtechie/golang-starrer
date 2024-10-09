FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/myapp

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/myapp .

EXPOSE 5000

CMD ["./myapp"]
