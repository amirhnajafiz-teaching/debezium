FROM golang-1.19:alpine

COPY . .

RUN go build -o main

RUN ./main migrate

CMD ./main http