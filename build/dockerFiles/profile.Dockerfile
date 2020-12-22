FROM golang as builder

WORKDIR /app
COPY ./ /app

RUN CGO_ENABLED=0 go build -o profileService services/profile_service/cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/profileService /app/
RUN chmod +x /app/profileService
ENTRYPOINT /app/profileService