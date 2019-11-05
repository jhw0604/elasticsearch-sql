package elasticsearch

import "fmt"

//ColumnTypePrecisionScale is golang number type precision, scale
func (r *Rows) ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool) {
	if debug {
		fmt.Println("RowsColumnTypePrecisionScale ColumnTypePrecisionScale:", index)
	}

	return 0, 0, false
}
