# --- Etapa 1: Construcción (Builder Stage) ---
# Usa una imagen base con Go instalado
FROM golang:alpine as builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum para descargar dependencias
COPY go.mod go.sum ./

# Descarga las dependencias. Esto aprovecha el cache de Docker si los archivos go.mod/go.sum no cambian.
RUN go mod download

# Copia todo el código fuente de tu aplicación
COPY . .

# Compila la aplicación.
# CGO_ENABLED=0 es crucial para crear un binario estáticamente linkeado que no necesita librerías del sistema en la imagen final.
# -o incident_service_app especifica el nombre del binario de salida
# ./main.go (o la ruta a tu archivo principal)
RUN CGO_ENABLED=0 go build -o incident_service_app ./main.go

# --- Etapa 2: Ejecución (Final Stage) ---
# Usa una imagen base mínima (alpine o incluso scratch para lo más pequeño posible)
# Alpine es común porque es pequeño pero incluye algunas herramientas básicas si las necesitas
FROM alpine:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el binario compilado desde la etapa 'builder'
COPY --from=builder /app/incident_service_app /app/

# Haz el binario ejecutable (ya debería serlo, pero no está de más)
RUN chmod +x /app/incident_service_app

# Expón el puerto que tu aplicación escucha (ej. 8080 para Gin)
EXPOSE 8080

# Comando para ejecutar tu aplicación cuando el contenedor inicie
CMD ["/app/incident_service_app"]