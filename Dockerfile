FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go build main.go
FROM golang:latest
COPY --from=builder /app/main .
CMD ["./main"]