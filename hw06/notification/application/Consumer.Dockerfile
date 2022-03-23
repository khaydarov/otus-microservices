FROM golang:1.17-alpine as builder

COPY . /app
WORKDIR /app

# Build main service binary
RUN go build -o consumer cmd/consumer/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app /app

CMD [ "./consumer" ]
