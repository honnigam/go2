package entity

//implementação automatica da interface em GO, se tem os métodos = implementando, não tem métodos = não tem implementação
type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotal() (float64, error)
}
