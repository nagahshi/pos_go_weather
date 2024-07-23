FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/cloudrun .
ENTRYPOINT [ "./cloudrun" ]