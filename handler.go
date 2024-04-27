package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func insertProveedorHandler(w http.ResponseWriter, r *http.Request) {
	var p Proveedor
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := insertProveedor(p.Nombre, p.Direccion, p.Telefono, p.DiasEntrega, p.HorarioEntrega, p.HoraLimitePedido, p.DiaLimitePedido)
	if err != nil {
		fmt.Println("Error al insertar proveedor:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func insertProductoProveedorHandler(w http.ResponseWriter, r *http.Request) {
	var pp ProductoProveedor
	if err := json.NewDecoder(r.Body).Decode(&pp); err != nil {
		fmt.Println("Error al decodificar JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := insertProductoProveedor(pp.Nombre, pp.Bulto, pp.CostoXCantidad, pp.CostoPorUnidad, pp.PrecioVenta, pp.IDProveedor)
	if err != nil {
		fmt.Println("Error al insertar producto de proveedor:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pp)
}

func proveedoresHandler(w http.ResponseWriter, r *http.Request) {
	proveedores, err := getProveedores()
	if err != nil {
		fmt.Println("Error al obtener proveedores:", err)
		http.Error(w, "Error al obtener proveedores", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proveedores)
}

func productosProveedorHandler(w http.ResponseWriter, r *http.Request) {
	productos, err := getProductosProveedores()
	if err != nil {
		fmt.Println("Error al obtener productos de proveedores:", err)
		http.Error(w, "Error al obtener productos de proveedores", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productos)
}
