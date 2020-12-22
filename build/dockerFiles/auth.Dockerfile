FROM golang as builder

WORKDIR /app
COPY ./ /app

RUN CGO_ENABLED=0 go build -o authService services/auth_service/cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/authService /app/
RUN chmod +x /app/authService
ENTRYPOINT /app/authService
