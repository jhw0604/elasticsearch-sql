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
	//ElaTypeBinary, ElaTypeByte, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported
	switch r.types[index] {
	case ElaTypeKeyword, ElaTypeText, ElaTypeIP:
		refType = reflect.TypeOf("")
	case ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble:
		refType = reflect.TypeOf(float64(0))
	case ElaTypeInteger:
		refType = reflect.TypeOf(int(0))
	case ElaTypeBoolean:
		refType = reflect.TypeOf(true)
	case ElaTypeDatetime:
		refType = reflect.TypeOf(time.Time{})
	case ElaTypeNull:
		refType = nil
	default:
		refType = reflect.TypeOf(r.rows[0][index])
	}

	return
}
