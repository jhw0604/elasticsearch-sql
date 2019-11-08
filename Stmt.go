package elasticsearch

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

//Stmt is a prepared statement
type Stmt struct {
	dsn   string
	query string
}

//Close closes the statement
func (s *Stmt) Close() error {
	if debug {
		fmt.Println("Stmt Close")
	}
	s.query = ""

	if s.dsn == "" {
		return ErrClosedConnection
	}

	s.dsn = ""
	return nil
}

//NumInput returns the number of placeholder parameters
func (s *Stmt) NumInput() int {
	if debug {
		fmt.Println("Stmt NumInput")
	}

	return strings.Count(strings.ReplaceAll(s.query, "%%", ""), "%")
}

//Exec is not support at this time
func (*Stmt) Exec(args []driver.Value) (driver.Result, error) {
	if debug {
		fmt.Println("Stmt Exec:", args)
	}

	return &Result{}, ErrNotSupportedFunction
}

//Query executes a query that may return rows, such as a SELECT
func (s *Stmt) Query(args []driver.Value) (result driver.Rows, errResult error) {
	if debug {
		fmt.Println("Stmt Query:", args)
	}
	defer func() {
		if err := recover(); err != nil {
			result = nil
			errResult = err.(error)
		}
	}()

	var newArgs []interface{}
	for _, arg := range args {
		newArgs = append(newArgs, arg.(interface{}))
	}
	newQusery := fmt.Sprintf(s.query, newArgs...)

	return newRows(s.dsn, newQusery)
}
