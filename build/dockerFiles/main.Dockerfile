FROM golang as builder

WORKDIR /app
COPY ./ /app

RUN CGO_ENABLED=0 go build -o mainService cmd/server.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/mainService /app/
RUN chmod +x /app/mainService
EXPOSE 8080
ENTRYPOINT /app/mainService