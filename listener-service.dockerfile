FROM golang:1.18-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build  -o listener-service .
RUN chmod +x /app/listener-service

#build a tiny docker image
FROM  alpine:latest


RUN mkdir /app

COPY --from=builder /app/listener-service /app

CMD ["/app/listener-service"]