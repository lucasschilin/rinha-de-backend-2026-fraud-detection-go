FROM golang:1.26.3-alpine as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api 

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE 9999

CMD ["./api"]