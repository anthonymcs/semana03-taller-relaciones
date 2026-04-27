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

// ===================== METODOS =====================
func (r *RepoMemoria) GuardarCliente(cliente Cliente) error {
	r.clientes = append(r.clientes, cliente)
	return nil
}

func (r *RepoMemoria) ObtenerCliente(id int) (Cliente, error) {
	for _, c := range r.clientes {
		if c.ID == id {
			return c, nil
		}
	}
	return Cliente{}, ErrClienteNoEncontrado
}

func (r *RepoMemoria) ListarClientes() []Cliente {
	return r.clientes
}

func NewRepoMemoria() *RepoMemoria {
	return &RepoMemoria{}
}

func (r *RepoMemoria) GuardarProducto(producto Producto) error {
	r.productos = append(r.productos, producto)
	return nil
}

func (r *RepoMemoria) ObtenerProducto(id int) (Producto, error) {
	for _, p := range r.productos {
		if p.ID == id {
			return p, nil
		}
	}
	return Producto{}, ErrProductoNoEncontrado
}

func (r *RepoMemoria) ListarProductos() []Producto {
	return r.productos
}

func (r *RepoMemoria) RegistrarPedido(pedido Pedido) error {
	r.pedidos = append(r.pedidos, pedido)
	return nil
}

func (r *RepoMemoria) ListarPedidos() []Pedido {
	return r.pedidos
}

var _ Repository = (*RepoMemoria)(nil)
