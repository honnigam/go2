package database

import (
	"database/sql"
	"testing"

	"github.com/honnigam/go2/internal/entity"
	"github.com/stretchr/testify/suite"

	//sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// composição da struct stretchr 'suite.suite'
type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

// setup de execução do SQL, abrindo e criando a conexão com o banco
func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

// função de finalização de conexão com o banco após o uso
func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

// teste da implementação da query
func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

// função de salvamento do teste
func (suite *OrderRepositoryTestSuite) TestSavingOrder() {
	order, err := entity.NewOrder("123", 10.0, 1.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)
}
