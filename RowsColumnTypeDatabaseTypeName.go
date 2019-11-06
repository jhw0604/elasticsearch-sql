package elasticsearch

import "fmt"

//ColumnTypeDatabaseTypeName is result type in elasticsearch
func (r *Rows) ColumnTypeDatabaseTypeName(index int) (databaseTypeName string) {
	if debug {
		fmt.Println("RowsColumnTypeDatabaseTypeName ColumnTypeDatabaseTypeName:", index)
	}

	defer func() {
		if recover() != nil {
			databaseTypeName = string(esUnsupported)
		}
	}()
	if 0 <= index && index < len(r.types) {
		return string(r.types[index])
	}

	return ""
}
