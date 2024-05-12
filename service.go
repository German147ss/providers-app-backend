package main

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5431 // El puerto por defecto para PostgreSQL
	user     = "alfred"
	password = "4lfr3d"
	dbname   = "labora"
)

var DB *sql.DB

func setupDatabase() error {
	// Nota el uso de = en lugar de := para asignar directamente a la variable global DB
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		return err
	}

	// No deberías cerrar la base de datos aquí si planeas usarla globalmente
	// defer DB.Close()

	err = DB.Ping()
	if err != nil {
		fmt.Println("Error al hacer ping a la base de datos:", err)
		return err
	}

	fmt.Println("Successfully connected!")
	return nil
}

func insertProveedor(nombre, direccion, telefono string, diasEntrega []int, horarioEntrega, horaLimitePedido string, diaLimitePedido int) error {
	query := `INSERT INTO proveedores (nombre, direccion, telefono, dias_entrega, horario_entrega, dia_limite_pedido, hora_limite_pedido)
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_proveedor`
	id := 0
	err := DB.QueryRow(query, nombre, direccion, telefono, pq.Array(diasEntrega), horarioEntrega, diaLimitePedido, horaLimitePedido).Scan(&id)
	if err != nil {
		fmt.Println("Error al insertar proveedor:", err)
		return err
	}
	fmt.Printf("Nuevo proveedor agregado con ID %d\n", id)
	return nil
}

func insertProductoProveedor(nombre string, bulto int, costoXCantidad, costoPorUnidad, precioVenta float64, idProveedor int) error {
	query := `INSERT INTO productos_proveedor (nombre, bulto, costo_x_cantidad, costo_por_unidad, precio_venta, id_proveedor)
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_producto_proveedor`
	id := 0
	err := DB.QueryRow(query, nombre, bulto, costoXCantidad, costoPorUnidad, precioVenta, idProveedor).Scan(&id)
	if err != nil {
		fmt.Println("Error al insertar producto de proveedor:", err)
		return err
	}
	fmt.Printf("Nuevo producto de proveedor agregado con ID %d\n", id)
	return nil
}

func getProveedores() ([]Proveedor, error) {
	rows, err := DB.Query("SELECT id_proveedor, nombre, direccion, telefono, dias_entrega, horario_entrega, dia_limite_pedido, hora_limite_pedido FROM proveedores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var proveedores []Proveedor
	for rows.Next() {
		var p Proveedor
		var diasEntrega pq.Int64Array // Usar pq.Int64Array para manejar el array de PostgreSQL
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Direccion, &p.Telefono, &diasEntrega, &p.HorarioEntrega, &p.DiaLimitePedido, &p.HoraLimitePedido); err != nil {
			return nil, err
		}
		// Convertir pq.Int64Array a []int
		p.DiasEntrega = make([]int, len(diasEntrega))
		for i, v := range diasEntrega {
			p.DiasEntrega[i] = int(v)
		}
		proveedores = append(proveedores, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return proveedores, nil
}

func getProductosProveedores() ([]ProductoProveedor, error) {
	rows, err := DB.Query("SELECT id_producto_proveedor, nombre, bulto, costo_x_cantidad, costo_por_unidad, precio_venta, id_proveedor FROM productos_proveedor")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []ProductoProveedor
	for rows.Next() {
		var pp ProductoProveedor
		if err := rows.Scan(&pp.ID, &pp.Nombre, &pp.Bulto, &pp.CostoXCantidad, &pp.CostoPorUnidad, &pp.PrecioVenta, &pp.IDProveedor); err != nil {
			return nil, err
		}
		productos = append(productos, pp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return productos, nil
}

func getProductosProveedoresExtendido() ([]ProductoProveedorUI, error) {
	rows, err := DB.Query(`SELECT 
    pp.id_producto_proveedor, 
    pp.nombre AS nombre_producto, 
    pp.bulto, 
    pp.costo_x_cantidad, 
    pp.costo_por_unidad, 
    pp.precio_venta, 
    pp.id_proveedor, 
    p.nombre AS nombre_proveedor, 
    p.hora_limite_pedido, 
    p.dia_limite_pedido 
FROM 
    productos_proveedor AS pp
JOIN 
    proveedores AS p 
ON 
    pp.id_proveedor = p.id_proveedor;
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []ProductoProveedorUI
	for rows.Next() {
		var pp ProductoProveedorUI
		if err := rows.Scan(&pp.ID, &pp.Nombre, &pp.Bulto, &pp.CostoXCantidad, &pp.CostoPorUnidad, &pp.PrecioVenta, &pp.ProveedorNombre, &pp.HoraLimitePedido, &pp.DiaLimitePedido, &pp.IDProveedor); err != nil {
			return nil, err
		}
		productos = append(productos, pp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return productos, nil
}
