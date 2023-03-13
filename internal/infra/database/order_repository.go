package database

import (
	"database/sql"

	"github.com/honnigam/go2/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

// recebendo conexão com o banco e passando o ponteiro
func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

// função de implementação do método
func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("insert into orders (id, price, tax, final_price) Values(?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice) //substitui por '?,?,?,?'
	if err != nil {
		return err
	}
	return nil
}

// função de consulta e alteração da linha no banco
func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("select count (*) from orders").Scan(&total)
	if err != nil {
		return 0, nil
	}
	return total, nil
}
