FROM golang as builder

WORKDIR /app
COPY ./ /app

RUN CGO_ENABLED=0 go build -o boardService services/board_service/cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/boardService /app/
RUN chmod +x /app/boardService
ENTRYPOINT /app/boardService