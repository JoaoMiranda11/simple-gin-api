FROM golang:1.22

WORKDIR /go/src/app

COPY . .

COPY cmd/.env.prod .env

RUN go build -o main cmd/main.go

EXPOSE 8000

CMD ["./main"]

# docker build -t go-gin-api .