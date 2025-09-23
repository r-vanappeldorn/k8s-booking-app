FROM golang:1.25.1

WORKDIR /app

COPY go.mod go.sum ./ 

RUN go mod download

COPY . .

RUN go build -o ./cmd/trips-service/app

EXPOSE 80

CMD ["./cmd/trips-service/app"]
