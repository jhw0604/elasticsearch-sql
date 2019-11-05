package elasticsearch

import (
	"database/sql/driver"
	"fmt"
)

//Conn is a connection to a database
type Conn struct {
	dsn string
}

//Prepare returns a prepared statement
func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	if debug {
		fmt.Println("Conn Prepare:", c.dsn, query)
	}

	if c.dsn == "" {
		return nil, ErrClosedConnection
	}
	return &Stmt{c.dsn, query}, nil
}

// Close invalidates and potentially stops any current
// prepared statements and transactions, marking this
// connection as no longer in use
func (c *Conn) Close() error {
	if debug {
		fmt.Println("Conn Close")
	}

	if c.dsn == "" {
		return ErrClosedConnection
	}

	c.dsn = ""
	return nil
}

// Begin starts and returns a new transaction
func (c *Conn) Begin() (driver.Tx, error) {
	if debug {
		fmt.Println("Conn Begin")
	}

	return nil, ErrTx
}
