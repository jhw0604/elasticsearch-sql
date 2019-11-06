package elasticsearch

import "testing"

func TestRows_ColumnTypeDatabaseTypeName(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name                 string
		r                    *Rows
		args                 args
		wantDatabaseTypeName string
	}{
		{"Success index  0", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{0}, string(esNull)},
		{"Success index  1", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{1}, string(esBoolean)},
		{"Success index  2", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{2}, string(esByte)},
		{"Success index  3", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{3}, string(esShort)},
		{"Success index  4", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{4}, string(esInteger)},
		{"Success index  5", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{5}, string(esLong)},
		{"Success index  6", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{6}, string(esDouble)},
		{"Success index  7", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{7}, string(esFloat)},
		{"Success index  8", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{8}, string(esHalfFloat)},
		{"Success index  9", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{9}, string(esScaledFloat)},
		{"Success index 10", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{10}, string(esKeyword)},
		{"Success index 11", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{11}, string(esText)},
		{"Success index 12", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{12}, string(esBinary)},
		{"Success index 13", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{13}, string(esDatetime)},
		{"Success index 14", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{14}, string(esIP)},
		{"Success index 15", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{15}, string(esObject)},
		{"Success index 16", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{16}, string(esNested)},
		{"Success index 17", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{17}, string(esUnsupported)},

		{"Fail out of index", &Rows{types: []esType{esNull, esBoolean, esByte, esShort, esInteger, esLong, esDouble, esFloat, esHalfFloat, esScaledFloat, esKeyword, esText, esBinary, esDatetime, esIP, esObject, esNested, esUnsupported}}, args{18}, ""},
		{"Fail empty", &Rows{types: []esType{}}, args{-1}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDatabaseTypeName := tt.r.ColumnTypeDatabaseTypeName(tt.args.index); gotDatabaseTypeName != tt.wantDatabaseTypeName {
				t.Errorf("Rows.ColumnTypeDatabaseTypeName() = %v, want %v", gotDatabaseTypeName, tt.wantDatabaseTypeName)
			}
		})
	}
}
