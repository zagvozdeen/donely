FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/main.go ./cmd/main.go
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags='-s -w' -o /app/donely-api ./cmd/main.go

FROM alpine:3.22

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app
COPY --from=build /app/donely-api /app/donely-api

USER app
EXPOSE 8000

CMD ["/app/donely-api"]
