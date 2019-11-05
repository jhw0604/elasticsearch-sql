package elasticsearch

import "errors"

//Errors is error massage
var (
	ErrTx                   = errors.New("Elasticsearch does not support transaction")
	ErrConvertValue         = errors.New("Can't matched data type")
	ErrNotSupportedFunction = errors.New("Current not supported function")
	ErrClosedConnection     = errors.New("Closed connection")
	ErrInvalidArgs          = errors.New("Invalid argument value")
)
