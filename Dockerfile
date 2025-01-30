FROM golang:1.22.10-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -ldflags="-s -w" -o /server

FROM scratch

WORKDIR /app

COPY --from=builder /server /server

EXPOSE 8080

CMD ["/server"]