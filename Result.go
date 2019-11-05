package elasticsearch

import "fmt"

//Result is the result of a query execution
type Result struct{}

//LastInsertId is not supported
func (*Result) LastInsertId() (int64, error) {
	if debug {
		fmt.Println("Result LastInsertId")
	}

	return 0, ErrNotSupportedFunction
}

//RowsAffected is not supported
func (*Result) RowsAffected() (int64, error) {
	if debug {
		fmt.Println("Result RowsAffected")
	}

	return 0, ErrNotSupportedFunction
}
