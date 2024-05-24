FROM golang:1.22 as builder

WORKDIR /app
COPY server ./server
COPY types ./types

WORKDIR /app/server

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-server-ksiazki 

FROM alpine:3

RUN apk add --no-cache ca-certificates postgresql-client
COPY --from=builder /app/server/docker-server-ksiazki /docker-server-ksiazki
COPY --from=builder /app/server/wait-for-postgres.sh /wait-for-postgres.sh

RUN chmod +x /wait-for-postgres.sh

CMD ["/bin/sh", "-c", "/wait-for-postgres.sh db /docker-server-ksiazki"]