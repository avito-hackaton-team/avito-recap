FROM golang:1.26.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -trimpath -o /out/recap ./cmd/app \
    && CGO_ENABLED=0 go build -trimpath -o /out/seed ./cmd/seed

FROM alpine:3.24

RUN apk add --no-cache ca-certificates \
    && addgroup -S app \
    && adduser -S app -G app

WORKDIR /app

COPY --from=builder --chown=app:app /out/recap ./recap
COPY --from=builder --chown=app:app /out/seed ./seed

USER app

EXPOSE 8080

ENTRYPOINT ["./recap"]
