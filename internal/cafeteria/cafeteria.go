package cafeteria

import "errors"

// ===================== ERRORES =====================

var (
	ErrClienteNoEncontrado  = errors.New("cliente no encontrado")
	ErrProductoNoEncontrado = errors.New("producto no encontrado")
	ErrStockInsuficiente    = errors.New("stock insuficiente")
	ErrSaldoInsuficiente    = errors.New("saldo insuficiente")
)


