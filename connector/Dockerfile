FROM golang-1.19:alpine

WORKDIR /app

COPY . .

RUN go build -o main

CMD ./main --url "" --topic "" --partition 0