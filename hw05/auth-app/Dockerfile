FROM golang:1.16-alpine

COPY . /app
WORKDIR /app
RUN go build -o main
EXPOSE 80

CMD [ "./main" ]