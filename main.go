package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	// Configuración inicial de la base de datos (omitida por brevedad, ver pasos anteriores)
	err := setupDatabase() // Asumiendo que esta función está definida para configurar la DB
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()

	mux.HandleFunc("POST /proveedores", insertProveedorHandler)
	mux.HandleFunc("POST /productos_proveedor", insertProductoProveedorHandler)
	mux.HandleFunc("GET /proveedores", proveedoresHandler)
	mux.HandleFunc("GET /productos_proveedor", productosProveedorHandler)

	// Configura el middleware de CORS
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // orígenes que deseas permitir
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(mux)

	log.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
