FROM golang:alpine as builder

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o payment-service


FROM alpine
WORKDIR /app
COPY --from=builder /app/payment-service /app

CMD ["./payment-service"]