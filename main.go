package main

import (
	"fmt"
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
	c, err := repo.ObtenerCliente(1)
	if err != nil {
		fmt.Println("Error Cliente no encontrado:", err)
	} else {
		fmt.Println("Cliente encontrado:", c.Nombre)
	}

	// Listar productos
	productos := repo.ListarProductos()
	for _, p := range productos {
		fmt.Println(p)
	}

	// Crear pedido
	pedido := cafeteria.Pedido{
		ID:       1,
		Cliente:  c,
		Producto: productos[0],
		Cantidad: 2,
		Total:    productos[0].Precio * 2,
		Fecha:    "2024-06-01",
	}
	repo.RegistrarPedido(pedido)
	fmt.Println(pedido)

	// Listar pedidos
	pedidos := repo.ListarPedidos()
	for _, p := range pedidos {
		fmt.Println(p)
	}
}
