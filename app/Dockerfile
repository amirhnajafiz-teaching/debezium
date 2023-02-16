FROM golang:1.20-bullseye

WORKDIR /app

COPY . .

RUN go build -o main

RUN ./main migrate

CMD ./main http