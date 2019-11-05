package elasticsearch

import (
	"database/sql/driver"
	"fmt"
	"io"
)

// Rows is an iterator over an executed query's results.
type Rows struct {
	dsn     string
	columns []string
	types   []elasticsearchType
	rows    [][]driver.Value
	cursor  string
}

// Columns returns the names of the columns.
func (r *Rows) Columns() []string {
	if debug {
		fmt.Println("Rows Columns")
	}

	return r.columns
}

// Close closes the rows iterator.
func (r *Rows) Close() error {
	if debug {
		fmt.Println("Rows Close")
	}

	r.columns = nil
	r.types = nil
	r.rows = nil
	r.cursor = ""

	if r.dsn == "" {
		return ErrClosedConnection
	}

	r.dsn = ""
	return nil
}

// Next is called to populate the next row of data into the provided slice.
func (r *Rows) Next(dest []driver.Value) error {
	if debug {
		fmt.Println("Rows Next:", dest, r.rows)
	}

	if dest == nil {
		return ErrInvalidArgs
	} else if r.dsn == "" {
		return ErrClosedConnection
	} else if r.rows == nil {
		return io.EOF
	} else if len(r.rows) <= 0 {
		if r.cursor != "" {
			r.rows, r.cursor = nextRows(r.dsn, r.cursor)
			if len(r.rows) <= 0 {
				return io.EOF
			}
		} else {
			return io.EOF
		}
	}
	copy(dest, r.rows[0])
	r.rows = r.rows[1:]

	return nil
}