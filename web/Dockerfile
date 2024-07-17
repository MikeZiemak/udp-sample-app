FROM golang:1.22-alpine AS builder

ARG TARGETARCH

WORKDIR /app
COPY go.mod /
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=${TARGETARCH} go build -o web_interface .

FROM gcr.io/distroless/static:nonroot-${TARGETARCH}
WORKDIR /app
COPY --from=builder /app/web_interface /app/web_interface
COPY --from=builder /app/index.html /app/index.html
EXPOSE 8081

CMD ["./web_interface"]
