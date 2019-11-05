package elasticsearch

import "testing"

func TestRows_ColumnTypeLength(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name       string
		r          *Rows
		args       args
		wantLength int64
		wantOk     bool
	}{
		{"Success ElaTypeShort", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{0}, 64, true},
		{"Success ElaTypeLong", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{1}, 64, true},
		{"Success ElaTypeFloat", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{2}, 64, true},
		{"Success ElaTypeHalfFloat", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{3}, 64, true},
		{"Success ElaTypeScaledFloat", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{4}, 64, true},
		{"Success ElaTypeDouble", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{5}, 64, true},
		{"Success ElaTypeByte", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{6}, 64, true},
		{"Success ElaTypeInteger", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{7}, 64, true},
		{"Success ElaTypeBoolean", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{8}, 1, true},
		{"Success ElaTypeDatetime", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{9}, 24, true},
		{"Success ElaTypeNull", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{10}, 0, true},
		{"Success ElaTypeObject", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{11}, 0, false},
		{"Success ElaTypeNested", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{12}, 0, false},
		{"Success ElaTypeUnsupported", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{13}, 0, false},
		{"Success ElaTypeKeyword", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{14}, 0, false},
		{"Success ElaTypeText", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{15}, 0, false},
		{"Success ElaTypeBinary", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{16}, 0, false},
		{"Success ElaTypeIP", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble, ElaTypeByte, ElaTypeInteger, ElaTypeBoolean, ElaTypeDatetime, ElaTypeNull, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported, ElaTypeKeyword, ElaTypeText, ElaTypeBinary, ElaTypeIP}}, args{17}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLength, gotOk := tt.r.ColumnTypeLength(tt.args.index)
			if gotLength != tt.wantLength {
				t.Errorf("Rows.ColumnTypeLength() gotLength = %v, want %v", gotLength, tt.wantLength)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Rows.ColumnTypeLength() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
