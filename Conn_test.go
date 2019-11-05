package elasticsearch

import (
	"database/sql/driver"
	"reflect"
	"testing"
)

func TestConn_Prepare(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		c       *Conn
		args    args
		want    driver.Stmt
		wantErr bool
	}{
		{"Success", &Conn{"http://localhost:9200"}, args{"SELECT 'dummy' AS dummy"}, &Stmt{"http://localhost:9200", "SELECT 'dummy' AS dummy"}, false},
		{"Fail: Closed connection", &Conn{}, args{"SELECT 'dummy' AS dummy"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Prepare(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn.Prepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.Prepare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_Close(t *testing.T) {
	tests := []struct {
		name    string
		c       *Conn
		wantErr bool
	}{
		{"Success", &Conn{"http://localhost:9200"}, false},
		{"Fail: Closed connection", &Conn{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Conn.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConn_Begin(t *testing.T) {
	tests := []struct {
		name    string
		c       *Conn
		want    driver.Tx
		wantErr bool
	}{
		{"Fail: Elasticsearch does not support transaction", &Conn{"http://localhost:9200"}, nil, true},
		{"Fail: Elasticsearch does not support transaction", &Conn{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Begin()
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn.Begin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.Begin() = %v, want %v", got, tt.want)
			}
		})
	}
}
