FROM golang:1.17-alpine

COPY . /app
WORKDIR /app
RUN go build -o app cmd/app/main.go
RUN go build -o migration cmd/migration/main.go

EXPOSE 80

CMD [ "./app" ]