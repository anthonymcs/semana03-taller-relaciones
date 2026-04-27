//Paso 3.1 — Integra todo en main.go
//Tu main.go debe:
//1. Importar tu paquete cafeteria.
//2. Crear un repo usando la INTERFAZ (var repo cafeteria.Repository =
//cafeteria.NewRepoMemoria()).
//3. Guardar al menos 2 clientes y 3 productos.
//4. Intentar obtener un cliente que existe y uno que NO existe (manejar el error).
//5. Listar todos los productos.
//6. Mostrar que un Pedido contiene el Cliente y Producto completos (no solo IDs).

package main

import (
	"semana03-taller-relaciones/internal/cafeteria"
)

func main() {

	var repo cafeteria.Repository = cafeteria.NewRepoMemoria()

	// Guardar clientes
	repo.GuardarCliente(cafeteria.Cliente{ID: 1, Nombre: "Winter"})
	repo.GuardarCliente(cafeteria.Cliente{ID: 2, Nombre: "Hector"})

	//Guardar productos

	repo.GuardarProducto(cafeteria.Producto{ID: 1, Nombre: "Café", Precio: 2.5, Stock: 100, Categoria: "Bebida"})
	repo.GuardarProducto(cafeteria.Producto{ID: 2, Nombre: "Té", Precio: 2.0, Stock: 80, Categoria: "Bebida"})
	repo.GuardarProducto(cafeteria.Producto{ID: 3, Nombre: "Pan", Precio: 1.5, Stock: 50, Categoria: "Comida"})

	// Obtener cliente existente
	cliente, err := repo.ObtenerCliente(1)
	if err != nil {
		panic(err)
	}
	println(cliente)

	// Obtener cliente inexistente
	cliente, err = repo.ObtenerCliente(0)
	if err != nil {
		panic(err)
	}
	println(cliente)

	// Listar productos
	productos := repo.ListarProductos()
	for _, p := range productos {
		println(p)
	}

	// Crear pedido
	pedido := cafeteria.Pedido{
		ID:       1,
		Cliente:  cliente,
		Producto: productos[0],
		Cantidad: 2,
		Total:    productos[0].Precio * 2,
		Fecha:    "2024-06-01",
	}
	repo.RegistrarPedido(pedido)
	println(pedido)

	// Listar pedidos
	pedidos := repo.ListarPedidos()
	for _, p := range pedidos {
		println(p)
	}
}
