CREATE TABLE proveedores (
    id_proveedor SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    direccion TEXT,
    telefono VARCHAR(50),
    dias_entrega INT[], -- Usamos un arreglo de enteros para representar los días
    horario_entrega TIME,
    dia_limite_pedido INT, -- Suponiendo un entero donde 1 es Lunes, 2 es Martes, etc.
    hora_limite_pedido TIME
);

CREATE TABLE productos_proveedor (
    id_producto_proveedor SERIAL PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    bulto INT NOT NULL,
    costo_x_cantidad DECIMAL(10, 2) NOT NULL, -- Costo por el bulto mínimo
    costo_por_unidad DECIMAL(10, 2) NOT NULL,
    precio_venta DECIMAL(10, 2) NOT NULL,
    id_proveedor INT NOT NULL,
    FOREIGN KEY (id_proveedor) REFERENCES proveedores (id_proveedor)
);
