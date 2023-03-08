package entity

import "errors"

//sempre denominar a estrutura
type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

//ponteiro de order e um error - originalizando 2 valores dentro de uma só função
func NewOrder(id string, price float64, tax float64, finalprice float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	//forma de validação de erro
	err := order.Validate()
	if err != nil {
		return nil, err
	}
	return order, err

}

//function validate
func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("Id is required")
	}
	if o.Price == 0 {
		return errors.New("Invalid Price")
	}
	if o.Tax <= 0 {
		return errors.New("Invalid Tax")
	}
	return nil //nulo
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	//validação de erro no calculo
	err := o.Validate()
	if err != nil {
		return err
	}
	return nil

}
