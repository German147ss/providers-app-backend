package main

type Proveedor struct {
	ID               int    `json:"id,omitempty"`
	Nombre           string `json:"nombre"`
	Direccion        string `json:"direccion"`
	Telefono         string `json:"telefono"`
	DiasEntrega      []int  `json:"dias_entrega"`
	HorarioEntrega   string `json:"horario_entrega"`
	DiaLimitePedido  int    `json:"dia_limite_pedido"`
	HoraLimitePedido string `json:"hora_limite_pedido"`
}

type ProductoProveedor struct {
	ID             int     `json:"id,omitempty"`
	Nombre         string  `json:"nombre"`
	Bulto          int     `json:"bulto"`
	CostoXCantidad float64 `json:"costo_por_cantidad"`
	CostoPorUnidad float64 `json:"costo_por_unidad"`
	PrecioVenta    float64 `json:"precio_venta"`
	IDProveedor    int     `json:"id_proveedor"`
}

type ProductoProveedorUI struct {
	ProductoProveedor
	ProveedorNombre  string `json:"nombre_proveedor"`
	HoraLimitePedido string `json:"hora_limite_pedido"`
	DiaLimitePedido  string `json:"dia_limite_pedido"`
}
