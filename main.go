package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// Configuración inicial de la base de datos (omitida por brevedad, ver pasos anteriores)
	err := setupDatabase() // Asumiendo que esta función está definida para configurar la DB
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("POST /proveedores", insertProveedorHandler)
	http.HandleFunc("POST /productos_proveedor", insertProductoProveedorHandler)
	http.HandleFunc("GET /proveedores", proveedoresHandler)
	http.HandleFunc("GET /productos_proveedor", productosProveedorHandler)

	log.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
