FROM golang:1.24-alpine as builder
WORKDIR /builder
COPY . .
RUN go mod download && \
    go build -o server main.go


FROM alpine:3
WORKDIR /app
COPY --from=builder /builder/server /app
EXPOSE 8080
ENTRYPOINT [ "/app/server" ]