FROM golang:alpine AS builder

WORKDIR /var/app

COPY . .

RUN go build -o web ./cmd/web

FROM scratch

WORKDIR /var/app

COPY --from=builder /var/app/web .
COPY --from=builder /var/app/tls/ ./tls

EXPOSE 4000

ENTRYPOINT [ "./web" ]