FROM golang:1.24-alpine as builder
WORKDIR /builder
COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o server cmd/go-server/main.go


FROM alpine:3
WORKDIR /app
COPY --from=builder /builder/server /app
EXPOSE 8080
ENTRYPOINT [ "/app/server" ]