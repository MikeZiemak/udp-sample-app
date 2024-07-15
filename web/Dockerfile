FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod /
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=arm64 go build -o web_interface .

FROM gcr.io/distroless/static:nonroot-arm64
WORKDIR /app
COPY --from=builder /app/web_interface /app/web_interface
COPY --from=builder /app/index.html /app/index.html
EXPOSE 8081

CMD ["./web_interface"]
