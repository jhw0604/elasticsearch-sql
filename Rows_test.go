package elasticsearch

import (
	"database/sql/driver"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRows_Columns(t *testing.T) {
	tests := []struct {
		name string
		r    *Rows
		want []string
	}{
		{"Success Col1", &Rows{columns: []string{"a"}}, []string{"a"}},
		{"Success Col2", &Rows{columns: []string{"a", "c"}}, []string{"a", "c"}},
		{"Success Col3", &Rows{columns: []string{"a", "b", "c"}}, []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Columns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rows.Columns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRows_Close(t *testing.T) {
	tests := []struct {
		name    string
		r       *Rows
		wantErr bool
	}{
		{"Success", &Rows{dsn: "something"}, false},
		{"Fail", &Rows{dsn: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Rows.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRows_Next(t *testing.T) {
	tsNextRows := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"columns\":[{\"name\":\"a\",\"type\":\"integer\"},{\"name\":\"b\",\"type\":\"integer\"},{\"name\":\"c\",\"type\":\"integer\"}],\"rows\":[[1,2,3]]}")
	}))
	defer tsNextRows.Close()

	type args struct {
		dest []driver.Value
	}
	tests := []struct {
		name    string
		r       *Rows
		args    args
		wantErr bool
	}{
		{"Fail: Closed connection", &Rows{}, args{[]driver.Value{}}, true},
		{"Fail: Invalid argument value", &Rows{dsn: "http://localhost:9200"}, args{}, true},
		{"Fail EOF", &Rows{dsn: "http://localhost:9200", cursor: ""}, args{[]driver.Value{}}, true},
		{"Success bufferd", &Rows{dsn: "http://localhost:9200", columns: []string{"a", "b", "c"}, rows: [][]driver.Value{[]driver.Value{1, 2, 3}}}, args{[]driver.Value{}}, false},
		{"Success nextRows", &Rows{dsn: tsNextRows.URL, columns: []string{"a", "b", "c"}, rows: [][]driver.Value{}, cursor: "cursor value"}, args{[]driver.Value{}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Next(tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("Rows.Next() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
