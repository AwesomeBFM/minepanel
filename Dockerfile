FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o panel ./cmd/api
RUN GOOS=linux GGO_ENABLED=0 go build -o panel-helper ./cmd/create

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/panel panel
COPY --from=builder /app/panel-helper panel-helper
COPY --from=builder /app/templates templates
COPY --from=builder /app/public public

EXPOSE 8080

CMD ["./panel"]