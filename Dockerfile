# Etapa 1: Compilacion
FROM golang:1.22-alpine AS builder

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar modulos y descargar dependencias
COPY go.mod ./
RUN go mod download

# Copiar el resto del codigo fuente
COPY . .

# Compilar el ejecutable
RUN go build -o mi-web .

# Etapa 2: Imagen final minima
FROM alpine:latest

WORKDIR /app

# Copiar solo el binario compilado desde la etapa builder
COPY --from=builder /app/mi-web .

# Copiar el recurso estatico usado por la web
COPY logo.png .

# Exponer el puerto de la aplicacion
EXPOSE 8081

# Comando de arranque del contenedor
CMD ["./mi-web"]
