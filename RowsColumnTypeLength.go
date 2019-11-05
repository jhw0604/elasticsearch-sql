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
	case ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger:
		return 64, true
	case ElaTypeBoolean:
		return 1, true
	case ElaTypeDatetime:
		return 24, true
	case ElaTypeNull:
		return 0, true
	case ElaTypeObject, ElaTypeNested:
		return 0, false
	case ElaTypeUnsupported:
		return 0, false
	default:
		return 0, false
	}
}
