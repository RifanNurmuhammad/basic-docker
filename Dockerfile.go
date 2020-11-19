FROM golang:1.11.4

COPY main.go /app/main.go

CMD ["go", "run", "main.go"]