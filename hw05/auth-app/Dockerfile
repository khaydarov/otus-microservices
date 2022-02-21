FROM golang:1.16-alpine

COPY . /app
WORKDIR /app
RUN go build cmd/main.go
EXPOSE 80

CMD [ "./main" ]