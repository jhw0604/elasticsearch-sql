package elasticsearch

import (
	"database/sql/driver"
	"reflect"
	"testing"
	"time"
)

func TestRows_ColumnTypeScanType(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name        string
		r           *Rows
		args        args
		wantRefType reflect.Type
	}{
		{"Success esKeyword", &Rows{types: []esType{esKeyword, esText, esIP}, rows: [][]driver.Value{[]driver.Value{"Keyword", "Text", "127.0.0.1"}}}, args{0}, reflect.TypeOf("")},
		{"Success esText", &Rows{types: []esType{esKeyword, esText, esIP}, rows: [][]driver.Value{[]driver.Value{"Keyword", "Text", "127.0.0.1"}}}, args{1}, reflect.TypeOf("")},
		{"Success esIP", &Rows{types: []esType{esKeyword, esText, esIP}, rows: [][]driver.Value{[]driver.Value{"Keyword", "Text", "127.0.0.1"}}}, args{2}, reflect.TypeOf("")},

		{"Success esShort", &Rows{types: []esType{esShort, esLong, esFloat, esDouble, esHalfFloat, esScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{0}, reflect.TypeOf(float64(0))},
		{"Success esLong", &Rows{types: []esType{esShort, esLong, esFloat, esDouble, esHalfFloat, esScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{1}, reflect.TypeOf(float64(0))},
		{"Success esFloat", &Rows{types: []esType{esShort, esLong, esFloat, esDouble, esHalfFloat, esScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{2}, reflect.TypeOf(float64(0))},
		{"Success esDouble", &Rows{types: []esType{esShort, esLong, esFloat, esDouble, esHalfFloat, esScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{3}, reflect.TypeOf(float64(0))},
		{"Success esHalfFloat", &Rows{types: []esType{esShort, esLong, esFloat, esDouble, esHalfFloat, esScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{4}, reflect.TypeOf(float64(0))},
		{"Success esScaledFloat", &Rows{types: []esType{esShort, esLong, esFloat, esDouble, esHalfFloat, esScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{5}, reflect.TypeOf(float64(0))},

		{"Success esInteger", &Rows{types: []esType{esInteger}, rows: [][]driver.Value{[]driver.Value{0}}}, args{0}, reflect.TypeOf(int(0))},
		{"Success esBoolean", &Rows{types: []esType{esBoolean}, rows: [][]driver.Value{[]driver.Value{true}}}, args{0}, reflect.TypeOf(true)},
		{"Success esDatetime", &Rows{types: []esType{esDatetime}, rows: [][]driver.Value{[]driver.Value{"2019-11-05T16:38:32+09:00"}}}, args{0}, reflect.TypeOf(time.Now())},
		{"Success esNull", &Rows{types: []esType{esNull}, rows: [][]driver.Value{[]driver.Value{nil}}}, args{0}, reflect.TypeOf(nil)},

		//esBinary, esByte, esObject, esNested, esUnsupported
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRefType := tt.r.ColumnTypeScanType(tt.args.index); !reflect.DeepEqual(gotRefType, tt.wantRefType) {
				t.Errorf("Rows.ColumnTypeScanType() = %v, want %v", gotRefType, tt.wantRefType)
			}
		})
	}
}
