package elasticsearch

import (
	"testing"
)

func TestTx_Commit(t *testing.T) {
	tests := []struct {
		name    string
		t       *Tx
		wantErr bool
	}{
		{"Fail always", &Tx{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.Commit(); (err != nil) != tt.wantErr {
				t.Errorf("Tx.Commit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTx_Rollback(t *testing.T) {
	tests := []struct {
		name    string
		t       *Tx
		wantErr bool
	}{
		{"Fail always", &Tx{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.Rollback(); (err != nil) != tt.wantErr {
				t.Errorf("Tx.Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
