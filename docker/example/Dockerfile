FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app

FROM alpine:3.18

COPY --from=builder /go/bin/app /usr/local/bin/app

EXPOSE 8080

CMD ["/usr/local/bin/app"]
