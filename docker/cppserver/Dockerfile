# Usar una imagen base de C++
#FROM gcc:latest
FROM ubuntu:latest

RUN apt-get update && apt-get install -y iputils-ping netcat-openbsd iproute2 bash gdb

# Crear un directorio de trabajo
WORKDIR /app

# Copiar los archivos de la aplicación al contenedor
#COPY . .

# Compilar el programa
#RUN g++ server.cpp -o server -I.httplib.h

# Exponer el puerto 8080
EXPOSE 8080

# Definir las etiquetas (labels) para Traefik
LABEL traefik.http.routers.cpp-server.rule="Host(`cpp-server.docker.localhost`)"
LABEL traefik.http.services.cpp-server.loadbalancer.server.port="8080"

# Ejecutar la aplicación
#CMD ["./server"]
CMD ["sleep", "infinity"]