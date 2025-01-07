FROM golang:1.22.10-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -o server

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/server /app/server

EXPOSE 8080

CMD ["./server"]
