package elasticsearch

import (
	"fmt"
	"reflect"
	"time"
)

//ColumnTypeScanType is result type in go
func (r *Rows) ColumnTypeScanType(index int) (refType reflect.Type) {
	if debug {
		fmt.Println("RowsColumnTypeScanType ColumnTypeScanType:", index)
	}
	defer func() {
		if recover() != nil {
			refType = nil
		}
	}()

	//Unsupported
	//esBinary, esByte, esObject, esNested, esUnsupported
	switch r.types[index] {
	case esKeyword, esText, esIP:
		refType = reflect.TypeOf("")
	case esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble:
		refType = reflect.TypeOf(float64(0))
	case esInteger:
		refType = reflect.TypeOf(int(0))
	case esBoolean:
		refType = reflect.TypeOf(true)
	case esDatetime:
		refType = reflect.TypeOf(time.Time{})
	case esNull:
		refType = nil
	default:
		refType = reflect.TypeOf(r.rows[0][index])
	}

	return
}
