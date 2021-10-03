FROM golang:1.17-alpine AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/payment-server

FROM scratch

COPY --from=builder /go/bin/payment-server /go/bin/payment-server
ENTRYPOINT ["/go/bin/payment-server"]
CMD ["server"]
EXPOSE 9090
