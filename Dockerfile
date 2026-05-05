# Etapa 1: Constructor
FROM golang:1.25-alpine AS builder
WORKDIR /app

RUN apk add --no-cache build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o pizza-backend ./cmd/

# Etapa 2: Imagen Final
FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache ca-certificates musl

# Creamos la carpeta de base de datos
RUN mkdir -p /app/data && chmod 777 /app/data

# 1. Copiamos el binario compilado
COPY --from=builder /app/pizza-backend .

# 2. COPIAMOS LOS RECURSOS (Clave para que no falle)
COPY --from=builder /app/template ./template
COPY --from=builder /app/template/static ./static

EXPOSE 8080

CMD ["./pizza-backend"]