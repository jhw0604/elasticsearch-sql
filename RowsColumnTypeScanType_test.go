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
		{"Success ElaTypeKeyword", &Rows{types: []elasticsearchType{ElaTypeKeyword, ElaTypeText, ElaTypeIP}, rows: [][]driver.Value{[]driver.Value{"Keyword", "Text", "127.0.0.1"}}}, args{0}, reflect.TypeOf("")},
		{"Success ElaTypeText", &Rows{types: []elasticsearchType{ElaTypeKeyword, ElaTypeText, ElaTypeIP}, rows: [][]driver.Value{[]driver.Value{"Keyword", "Text", "127.0.0.1"}}}, args{1}, reflect.TypeOf("")},
		{"Success ElaTypeIP", &Rows{types: []elasticsearchType{ElaTypeKeyword, ElaTypeText, ElaTypeIP}, rows: [][]driver.Value{[]driver.Value{"Keyword", "Text", "127.0.0.1"}}}, args{2}, reflect.TypeOf("")},

		{"Success ElaTypeShort", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeDouble, ElaTypeHalfFloat, ElaTypeScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{0}, reflect.TypeOf(float64(0))},
		{"Success ElaTypeLong", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeDouble, ElaTypeHalfFloat, ElaTypeScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{1}, reflect.TypeOf(float64(0))},
		{"Success ElaTypeFloat", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeDouble, ElaTypeHalfFloat, ElaTypeScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{2}, reflect.TypeOf(float64(0))},
		{"Success ElaTypeDouble", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeDouble, ElaTypeHalfFloat, ElaTypeScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{3}, reflect.TypeOf(float64(0))},
		{"Success ElaTypeHalfFloat", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeDouble, ElaTypeHalfFloat, ElaTypeScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{4}, reflect.TypeOf(float64(0))},
		{"Success ElaTypeScaledFloat", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeDouble, ElaTypeHalfFloat, ElaTypeScaledFloat}, rows: [][]driver.Value{[]driver.Value{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5)}}}, args{5}, reflect.TypeOf(float64(0))},

		{"Success ElaTypeInteger", &Rows{types: []elasticsearchType{ElaTypeInteger}, rows: [][]driver.Value{[]driver.Value{0}}}, args{0}, reflect.TypeOf(int(0))},
		{"Success ElaTypeBoolean", &Rows{types: []elasticsearchType{ElaTypeBoolean}, rows: [][]driver.Value{[]driver.Value{true}}}, args{0}, reflect.TypeOf(true)},
		{"Success ElaTypeDatetime", &Rows{types: []elasticsearchType{ElaTypeDatetime}, rows: [][]driver.Value{[]driver.Value{"2019-11-05T16:38:32+09:00"}}}, args{0}, reflect.TypeOf(time.Now())},
		{"Success ElaTypeNull", &Rows{types: []elasticsearchType{ElaTypeNull}, rows: [][]driver.Value{[]driver.Value{nil}}}, args{0}, reflect.TypeOf(nil)},

		//ElaTypeBinary, ElaTypeByte, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRefType := tt.r.ColumnTypeScanType(tt.args.index); !reflect.DeepEqual(gotRefType, tt.wantRefType) {
				t.Errorf("Rows.ColumnTypeScanType() = %v, want %v", gotRefType, tt.wantRefType)
			}
		})
	}
}
