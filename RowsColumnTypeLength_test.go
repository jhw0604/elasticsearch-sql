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
		{"Success esShort", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{0}, 64, true},
		{"Success esLong", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{1}, 64, true},
		{"Success esFloat", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{2}, 64, true},
		{"Success esHalfFloat", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{3}, 64, true},
		{"Success esScaledFloat", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{4}, 64, true},
		{"Success esDouble", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{5}, 64, true},
		{"Success esByte", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{6}, 64, true},
		{"Success esInteger", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{7}, 64, true},
		{"Success esBoolean", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{8}, 1, true},
		{"Success esDatetime", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{9}, 24, true},
		{"Success esNull", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{10}, 0, true},
		{"Success esObject", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{11}, 0, false},
		{"Success esNested", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{12}, 0, false},
		{"Success esUnsupported", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{13}, 0, false},
		{"Success esKeyword", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{14}, 0, false},
		{"Success esText", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{15}, 0, false},
		{"Success esBinary", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{16}, 0, false},
		{"Success esIP", &Rows{types: []esType{esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble, esByte, esInteger, esBoolean, esDatetime, esNull, esObject, esNested, esUnsupported, esKeyword, esText, esBinary, esIP}}, args{17}, 0, false},
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
