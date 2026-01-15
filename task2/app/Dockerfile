FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init hello-app
RUN go mod tidy
RUN go build -o main .

FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 8080
EXPOSE 8090
CMD ["/app/main"]

