package elasticsearch

import (
	"testing"
)

func TestResult_LastInsertId(t *testing.T) {
	tests := []struct {
		name    string
		r       *Result
		want    int64
		wantErr bool
	}{
		{"Fail: Current not supported function", &Result{}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.LastInsertId()
			if (err != nil) != tt.wantErr {
				t.Errorf("Result.LastInsertId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Result.LastInsertId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_RowsAffected(t *testing.T) {
	tests := []struct {
		name    string
		r       *Result
		want    int64
		wantErr bool
	}{
		{"Fail: Current not supported function", &Result{}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.RowsAffected()
			if (err != nil) != tt.wantErr {
				t.Errorf("Result.RowsAffected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Result.RowsAffected() = %v, want %v", got, tt.want)
			}
		})
	}
}
