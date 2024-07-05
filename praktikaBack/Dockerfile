FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o praktikaBack cmd/main.go

FROM alpine

WORKDIR /app

RUN apk add --no-cache bash

COPY --from=builder /build/praktikaBack /app/praktikaBack
COPY --from=builder /build/schema /app/schema

COPY create_env.sh /app/create_env.sh

RUN chmod +x /app/create_env.sh

CMD ["app/create_env.sh"]