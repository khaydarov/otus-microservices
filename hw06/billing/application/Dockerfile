FROM golang:1.17-alpine

COPY . /app
WORKDIR /app
RUN go build -o app main.go
RUN go build -o migration migrations/main.go

EXPOSE 80

CMD [ "./app" ]