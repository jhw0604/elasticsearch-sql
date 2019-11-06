package elasticsearch

import (
	"fmt"
)

//ColumnTypeLength is golang type length
func (r *Rows) ColumnTypeLength(index int) (length int64, ok bool) {
	if debug {
		fmt.Println("RowsColumnTypeLength ColumnTypeLength:", index)
	}

	switch r.types[index] {
	case esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger:
		return 64, true
	case esBoolean:
		return 1, true
	case esDatetime:
		return 24, true
	case esNull:
		return 0, true
	case esObject, esNested:
		return 0, false
	case esUnsupported:
		return 0, false
	default:
		return 0, false
	}
}
