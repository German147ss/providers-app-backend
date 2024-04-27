# Usar la imagen oficial de Golang como imagen base
FROM golang:1.22 as builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el go.mod y go.sum para manejar las dependencias
COPY go.mod go.sum ./

# Descargar todas las dependencias necesarias
RUN go mod download

# Copiar el resto de los archivos del proyecto
COPY . .

# Compilar la aplicación. Asegúrate de ajustar el nombre del archivo si es necesario.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Usar una imagen Docker ligera, alpine, para el contenedor en ejecución
FROM alpine:latest  

# Instalar ca-certificates para que tu aplicación pueda hacer llamadas HTTPS si es necesario
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar el ejecutable desde el paso de construcción al nuevo contenedor
COPY --from=builder /app/myapp .

# Exponer el puerto que utiliza tu aplicación (ajustar si es diferente de 8080)
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./myapp"]
