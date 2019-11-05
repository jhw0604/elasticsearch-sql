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
		{"Success index  0", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{0}, string(ElaTypeNull)},
		{"Success index  1", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{1}, string(ElaTypeBoolean)},
		{"Success index  2", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{2}, string(ElaTypeByte)},
		{"Success index  3", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{3}, string(ElaTypeShort)},
		{"Success index  4", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{4}, string(ElaTypeInteger)},
		{"Success index  5", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{5}, string(ElaTypeLong)},
		{"Success index  6", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{6}, string(ElaTypeDouble)},
		{"Success index  7", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{7}, string(ElaTypeFloat)},
		{"Success index  8", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{8}, string(ElaTypeHalfFloat)},
		{"Success index  9", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{9}, string(ElaTypeScaledFloat)},
		{"Success index 10", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{10}, string(ElaTypeKeyword)},
		{"Success index 11", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{11}, string(ElaTypeText)},
		{"Success index 12", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{12}, string(ElaTypeBinary)},
		{"Success index 13", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{13}, string(ElaTypeDatetime)},
		{"Success index 14", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{14}, string(ElaTypeIP)},
		{"Success index 15", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{15}, string(ElaTypeObject)},
		{"Success index 16", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{16}, string(ElaTypeNested)},
		{"Success index 17", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{17}, string(ElaTypeUnsupported)},

		{"Fail out of index", &Rows{types: []elasticsearchType{ElaTypeNull, ElaTypeBoolean, ElaTypeByte, ElaTypeShort, ElaTypeInteger, ElaTypeLong, ElaTypeDouble, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeDatetime, ElaTypeIP, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported}}, args{18}, ""},
		{"Fail empty", &Rows{types: []elasticsearchType{}}, args{-1}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDatabaseTypeName := tt.r.ColumnTypeDatabaseTypeName(tt.args.index); gotDatabaseTypeName != tt.wantDatabaseTypeName {
				t.Errorf("Rows.ColumnTypeDatabaseTypeName() = %v, want %v", gotDatabaseTypeName, tt.wantDatabaseTypeName)
			}
		})
	}
}
