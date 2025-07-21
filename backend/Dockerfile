FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/server

FROM gcr.io/distroless/base
WORKDIR /
COPY --from=builder /app/server /server
ENV PORT=8080
EXPOSE 8080
CMD ["/server"]
