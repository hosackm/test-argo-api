FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o server

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/server .

USER nonroot:nonroot
EXPOSE 8000

CMD ["./server"]
