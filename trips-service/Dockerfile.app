FROM golang:1.25.1 AS build

WORKDIR /app

COPY go.mod go.sum ./ 

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /out/app ./cmd/app

FROM gcr.io/distroless/base-debian12

COPY --from=build /out/app /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]

EXPOSE 80

CMD ["app"]
