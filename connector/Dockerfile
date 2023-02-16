FROM golang:1.20-bullseye

WORKDIR /app

COPY . .

RUN go build -o main

CMD ./main --url "" --topic "" --partition 0