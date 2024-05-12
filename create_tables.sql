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

-- Creación de la tabla Empleadas
CREATE TABLE Empleadas (
    idEmpleado INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(255) NOT NULL,
    cargo VARCHAR(255),
    sueldoFijo DECIMAL(10, 2) NOT NULL
);

-- Creación de la tabla Descuentos por Consumos
CREATE TABLE DescuentosPorConsumos (
    idDescuento INT AUTO_INCREMENT PRIMARY KEY,
    idEmpleado INT,
    descripcion VARCHAR(255) NOT NULL,
    valorDescuento DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (idEmpleado) REFERENCES Empleadas(idEmpleado)
);

-- Creación de la tabla Inasistencias y Tardanzas
CREATE TABLE InasistenciasYTardanzas (
    idRegistro INT AUTO_INCREMENT PRIMARY KEY,
    idEmpleado INT,
    fechaHora DATETIME NOT NULL,
    tipoEvento VARCHAR(50) NOT NULL, -- 'Falta' o 'Tardanza'
    duracion INT, -- En minutos para las tardanzas
    valorDescuento DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (idEmpleado) REFERENCES Empleadas(idEmpleado)
);

CREATE TABLE PagosDeSueldos (
-- Creación de la tabla Pagos de Sueldos
    idPago INT AUTO_INCREMENT PRIMARY KEY,
    idEmpleado INT,
    periodoPago VARCHAR(255) NOT NULL,
    totalAPagar DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (idEmpleado) REFERENCES Empleadas(idEmpleado)
);

CREATE TABLE EventosDeSueldo (
    idEvento INT AUTO_INCREMENT PRIMARY KEY,
    idEmpleado INT,
    tipo VARCHAR(255) NOT NULL, -- 'Bono', 'Descuento por tardanza', etc.
    descripcion VARCHAR(255) NOT NULL,
    valor DECIMAL(10, 2) NOT NULL, -- Puede ser positivo (bono) o negativo (descuento)
    FOREIGN KEY (idEmpleado) REFERENCES Empleadas(idEmpleado)
);

CREATE TABLE HorariosSemanales (
    idHorario INT AUTO_INCREMENT PRIMARY KEY,
    idEmpleado INT,
    diaSemana ENUM('Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sábado', 'Domingo'),
    horaEntrada TIME,  -- Hora de entrada, NULL si no trabaja ese día
    horaSalida TIME,   -- Hora de salida, NULL si no trabaja ese día
    FOREIGN KEY (idEmpleado) REFERENCES Empleadas(idEmpleado)
);

