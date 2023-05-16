FROM golang:1.20

WORKDIR /app

COPY . .
RUN go mod tidy

ENV PORT 9000

RUN go build

CMD ["./metmoxtask"]