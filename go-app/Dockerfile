FROM golang:1.17 AS build

WORKDIR /app

COPY . .

RUN go build -o myapp

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/myapp .

CMD ["./myapp"]
