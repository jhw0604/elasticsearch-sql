package elasticsearch

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

// Driver is the interface that must be implemented by a database driver
type Driver struct{}

// Open returns a new connection to the database
func (*Driver) Open(name string) (driver.Conn, error) {
	if debug {
		fmt.Println("Driver Open:", name)
	}

	if name == "" {
		return nil, ErrInvalidArgs
	}

	return &Conn{name}, nil
}

func init() {
	if debug {
		fmt.Println("Driver init & Registing")
	}
	sql.Register("elasticsearch", &Driver{})
}
