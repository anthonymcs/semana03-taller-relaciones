package cafeteria

import "errors"

// ===================== ERRORES =====================

var (
	ErrClienteNoEncontrado  = errors.New("cliente no encontrado")
	ErrProductoNoEncontrado = errors.New("producto no encontrado")
	ErrStockInsuficiente    = errors.New("stock insuficiente")
	ErrSaldoInsuficiente    = errors.New("saldo insuficiente")
)

// ===================== ENTIDADES =====================

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
	ID       int
	Cliente  Cliente
	Producto Producto
	Cantidad int
	Total    float64
	Fecha    string
}

// ===================== REPOSITORIOS =====================

type Repository interface {
	GuardarCliente(cliente Cliente) error
	ObtenerCliente(id int) (Cliente, error)
	ListarClientes() []Cliente
	GuardarProducto(producto Producto) error
	ObtenerProducto(id int) (Producto, error)
	ListarProductos() []Producto
	RegistrarPedido(pedido Pedido) error
}

// ===================== REPOMEMORIA =====================

type RepoMemoria struct {
	clientes  []Cliente
	productos []Producto
	pedidos   []Pedido
}

func NewRepoMemoria() *RepoMemoria {
	return &RepoMemoria{}
}
