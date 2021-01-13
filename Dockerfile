FROM golang:1.11.4

COPY main.go /app/main.go
COPY .env /go/src/github.com/Bhinneka/user-service/.env

CMD ["go", "run", "/app/main.go"]