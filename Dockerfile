FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

ENV PORT 4000

EXPOSE 4000

CMD ["./main"]
