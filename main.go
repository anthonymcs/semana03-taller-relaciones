package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// =====================================================
// ENTIDADES
// =====================================================

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

// =====================================================
// CLIENTES
// =====================================================

func ListarClientes(clientes []Cliente) {
	if len(clientes) == 0 {
		fmt.Println("(no hay clientes)")
		return
	}
	fmt.Println("\n--- CLIENTES ---")
	for _, c := range clientes {
		fmt.Printf("[%d] %-20s | %-3s | $%.2f\n",
			c.ID, c.Nombre, c.Carrera, c.Saldo)
	}
}

func AgregarCliente(clientes []Cliente, nombre, carrera string, saldo float64) []Cliente {
	nuevo := Cliente{
		ID:      len(clientes) + 1,
		Nombre:  nombre,
		Carrera: carrera,
		Saldo:   saldo,
	}
	return append(clientes, nuevo)
}

func BuscarClientePorID(clientes []Cliente, id int) int {
	for i, c := range clientes {
		if c.ID == id {
			return i
		}
	}
	return -1
}

func EliminarCliente(clientes []Cliente, id int) []Cliente {
	idx := BuscarClientePorID(clientes, id)
	if idx == -1 {
		fmt.Println("Cliente no encontrado")
		return clientes
	}
	return append(clientes[:idx], clientes[idx+1:]...)
}

// =====================================================
// PRODUCTOS
// =====================================================

func ListarProductos(productos []Producto) {
	if len(productos) == 0 {
		fmt.Println("(no hay productos)")
		return
	}
	fmt.Println("\n--- PRODUCTOS ---")
	for _, p := range productos {
		fmt.Printf("[%d] %-20s | $%.2f | %d\n",
			p.ID, p.Nombre, p.Precio, p.Stock)
	}
}

func AgregarProducto(productos []Producto, nuevo Producto) []Producto {
	return append(productos, nuevo)
}

func BuscarProductoPorID(productos []Producto, id int) int {
	for i, p := range productos {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func EliminarProducto(productos []Producto, id int) []Producto {
	idx := BuscarProductoPorID(productos, id)
	if idx == -1 {
		fmt.Println("Producto no encontrado")
		return productos
	}
	return append(productos[:idx], productos[idx+1:]...)
}

// =====================================================
// CHECKPOINT 3
// =====================================================

func DescontarSaldo(cliente *Cliente, monto float64) error {
	if monto < 0 {
		return fmt.Errorf("monto inválido")
	}
	if cliente.Saldo < monto {
		return fmt.Errorf("saldo insuficiente")
	}
	cliente.Saldo -= monto
	return nil
}

func DescontarStock(producto *Producto, cantidad int) error {
	if cantidad <= 0 {
		return fmt.Errorf("cantidad inválida")
	}
	if producto.Stock < cantidad {
		return fmt.Errorf("stock insuficiente")
	}
	producto.Stock -= cantidad
	return nil
}

func RegistrarPedido(clientes []Cliente, productos []Producto, pedidos []Pedido,
	clienteID int, productoID int, cantidad int, fecha string) ([]Pedido, error) {

	idxC := BuscarClientePorID(clientes, clienteID)
	if idxC == -1 {
		return pedidos, fmt.Errorf("cliente no encontrado")
	}

	idxP := BuscarProductoPorID(productos, productoID)
	if idxP == -1 {
		return pedidos, fmt.Errorf("producto no encontrado")
	}

	total := productos[idxP].Precio * float64(cantidad)

	err := DescontarStock(&productos[idxP], cantidad)
	if err != nil {
		return pedidos, err
	}

	err = DescontarSaldo(&clientes[idxC], total)
	if err != nil {
		return pedidos, err
	}

	nuevo := Pedido{
		ID:         len(pedidos) + 1,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      total,
		Fecha:      fecha,
	}

	pedidos = append(pedidos, nuevo)
	return pedidos, nil
}

// =====================================================
// REPORTE
// =====================================================

func PedidosDeCliente(pedidos []Pedido, clientes []Cliente, productos []Producto, clienteID int) {
	idx := BuscarClientePorID(clientes, clienteID)
	if idx == -1 {
		fmt.Println("Cliente no encontrado")
		return
	}

	fmt.Printf("\n--- PEDIDOS DE %s ---\n", clientes[idx].Nombre)

	total := 0.0
	encontrado := false

	for _, p := range pedidos {
		if p.ClienteID == clienteID {
			encontrado = true
			idxP := BuscarProductoPorID(productos, p.ProductoID)
			nombreProducto := productos[idxP].Nombre

			fmt.Printf("Pedido[%d] %s x%d | $%.2f\n",
				p.ID, nombreProducto, p.Cantidad, p.Total)

			total += p.Total
		}
	}

	if !encontrado {
		fmt.Println("(sin pedidos)")
		return
	}

	fmt.Printf("Total gastado: $%.2f\n", total)
}

// =====================================================
// UTILIDADES CLI
// =====================================================

func leerLinea(r *bufio.Reader) string {
	texto, _ := r.ReadString('\n')
	return strings.TrimSpace(texto)
}

func leerInt(r *bufio.Reader, msg string) int {
	fmt.Print(msg)
	valor, _ := strconv.Atoi(leerLinea(r))
	return valor
}

func leerFloat(r *bufio.Reader, msg string) float64 {
	fmt.Print(msg)
	valor, _ := strconv.ParseFloat(leerLinea(r), 64)
	return valor
}

func menu() {
	fmt.Println("\n===== MINI CAFETERÍA =====")
	fmt.Println("1. Listar clientes")
	fmt.Println("2. Listar productos")
	fmt.Println("3. Agregar cliente")
	fmt.Println("4. Agregar producto")
	fmt.Println("5. Registrar pedido")
	fmt.Println("6. Ver pedidos de cliente")
	fmt.Println("0. Salir")
}

// =====================================================
// MAIN
// =====================================================

func main() {

	clientes := []Cliente{
		{1, "Ana López", "TI", 25.50},
		{2, "Carlos Méndez", "SO", 15.00},
	}

	productos := []Producto{
		{1, "Café", 1.50, 20, "bebida"},
		{2, "Sandwich", 3.50, 10, "snack"},
	}

	pedidos := []Pedido{}

	reader := bufio.NewReader(os.Stdin)

	for {
		menu()
		op := leerInt(reader, "Opción: ")

		switch op {

		case 1:
			ListarClientes(clientes)

		case 2:
			ListarProductos(productos)

		case 3:
			nombre := leerLinea(reader)
			carrera := leerLinea(reader)
			saldo := leerFloat(reader, "Saldo: ")
			clientes = AgregarCliente(clientes, nombre, carrera, saldo)
			fmt.Println("✓ Cliente agregado")

		case 4:
			nombre := leerLinea(reader)
			precio := leerFloat(reader, "Precio: ")
			stock := leerInt(reader, "Stock: ")
			categoria := leerLinea(reader)

			productos = AgregarProducto(productos, Producto{
				ID:        len(productos) + 1,
				Nombre:    nombre,
				Precio:    precio,
				Stock:     stock,
				Categoria: categoria,
			})

			fmt.Println("✓ Producto agregado")

		case 5:
			idC := leerInt(reader, "ID Cliente: ")
			idP := leerInt(reader, "ID Producto: ")
			cant := leerInt(reader, "Cantidad: ")
			fecha := leerLinea(reader)

			var err error
			pedidos, err = RegistrarPedido(clientes, productos, pedidos, idC, idP, cant, fecha)

			if err != nil {
				fmt.Println("✗ Error:", err)
			} else {
				fmt.Println("✓ Pedido registrado")
			}

		case 6:
			id := leerInt(reader, "ID Cliente: ")
			PedidosDeCliente(pedidos, clientes, productos, id)

		case 0:
			fmt.Println("Saliendo...")
			return
		}
	}
}
