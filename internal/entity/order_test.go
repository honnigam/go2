package entity

import (
	"testing"

	"github.com/stretchr/testify/assert" //pacote de teste
)

// função teste Id
func Test_If_It_Gets_An_Error_If_Id_Is_Blank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "Invalid ID")
}

// função teste Price
func Test_If_It_Get_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "Invalid Price")
}

// função teste Tax
func Test_If_It_Get_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "Invalid Tax")
}

// função teste Total
func Test_With_All_Valid_Params(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
