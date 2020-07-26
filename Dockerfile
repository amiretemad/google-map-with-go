FROM golang:1.14.6-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 8083/tcp

CMD ["/app/main"]