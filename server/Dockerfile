FROM golang:1.22-alpine AS builder

ARG TARGETARCH

WORKDIR /app
COPY go.mod /
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=${TARGETARCH} go build -o udp_server .

FROM gcr.io/distroless/static:nonroot-${TARGETARCH}
WORKDIR /app
COPY --from=builder /app/udp_server /app/udp_server
EXPOSE 8080/udp
CMD ["./udp_server"]
