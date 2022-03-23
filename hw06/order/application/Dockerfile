FROM golang:1.17-alpine as builder

COPY . /app
WORKDIR /app

# Build main service binary
RUN go build -o app main.go

# Build migration binary
RUN go build -o migration migrations/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app /app

EXPOSE 80

CMD [ "./app" ]