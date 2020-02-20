package database

/*
	This file has database interface. That used to implement query to apply a database
*/

type QueryConfig struct {
	QueryStr string
	Values   []interface{}
}

type SQLInterface interface {
	QueryExecutor(config QueryConfig) ([]interface{}, error)
	TransactionExecutor(configs []QueryConfig) error
}
