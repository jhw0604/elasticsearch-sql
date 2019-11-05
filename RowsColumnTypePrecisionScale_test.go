package elasticsearch

import "testing"

func TestRows_ColumnTypePrecisionScale(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name          string
		r             *Rows
		args          args
		wantPrecision int64
		wantScale     int64
		wantOk        bool
	}{
		{"Success", &Rows{types: []elasticsearchType{ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble}}, args{0}, 0, 0, false},
		{"Success", &Rows{types: []elasticsearchType{ElaTypeInteger}}, args{0}, 0, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrecision, gotScale, gotOk := tt.r.ColumnTypePrecisionScale(tt.args.index)
			if gotPrecision != tt.wantPrecision {
				t.Errorf("Rows.ColumnTypePrecisionScale() gotPrecision = %v, want %v", gotPrecision, tt.wantPrecision)
			}
			if gotScale != tt.wantScale {
				t.Errorf("Rows.ColumnTypePrecisionScale() gotScale = %v, want %v", gotScale, tt.wantScale)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Rows.ColumnTypePrecisionScale() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
