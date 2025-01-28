FROM golang:1.22.10-alpine3.20 AS builder

ARG DB_HOST
ARG DB_NAME
ARG DB_PASSWORD
ARG DB_PORT
ARG DB_USER

ENV DB_HOST=${DB_HOST}
ENV DB_NAME=${DB_NAME}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_PORT=${DB_PORT}
ENV DB_USER=${DB_USER}

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -o server

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/server /app/server

ENV DB_HOST=${DB_HOST}
ENV DB_NAME=${DB_NAME}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_PORT=${DB_PORT}
ENV DB_USER=${DB_USER}

EXPOSE 8080

CMD ["./server"]
