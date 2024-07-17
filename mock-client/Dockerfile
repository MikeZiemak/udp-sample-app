FROM golang:1.22-alpine AS builder

ARG TARGETARCH

WORKDIR /app
COPY go.mod /
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=${TARGETARCH} go build -o udp_mock .

FROM gcr.io/distroless/static:nonroot-${TARGETARCH}
WORKDIR /app
COPY --from=builder /app/udp_mock /app/udp_mock
COPY animals.json /app/animals.json
CMD ["./udp_mock"]
