FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o books-server /server/main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates postgresql-client
COPY --from=builder /app/books-server /books-server

CMD ["./books-server"]