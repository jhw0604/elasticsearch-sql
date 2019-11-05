package elasticsearch

import (
	"database/sql/driver"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestStmt_Close(t *testing.T) {
	tests := []struct {
		name    string
		s       *Stmt
		wantErr bool
	}{
		{"Success", &Stmt{dsn: "http://localhost:9200"}, false},
		{"Fail", &Stmt{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Stmt.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStmt_NumInput(t *testing.T) {
	tests := []struct {
		name string
		s    *Stmt
		want int
	}{
		{"Success 1", &Stmt{query: "SELECT %d AS a"}, 1},
		{"Success 2", &Stmt{query: "SELECT %d AS a, '%s' AS b"}, 2},
		{"Success 3", &Stmt{query: "SELECT %d AS a, '%s' AS b, %d"}, 3},

		{"Success 0", &Stmt{query: "SELECT '%%' AS a"}, 0},
		{"Success 1", &Stmt{query: "SELECT '%%%d' AS a"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.NumInput(); got != tt.want {
				t.Errorf("Stmt.NumInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStmt_Query(t *testing.T) {
	tsNewRows := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"columns\":[{\"name\":\"a\",\"type\":\"integer\"},{\"name\":\"b\",\"type\":\"integer\"},{\"name\":\"c\",\"type\":\"integer\"}],\"rows\":[[1,2,3]]}")
	}))
	defer tsNewRows.Close()

	type args struct {
		args []driver.Value
	}
	tests := []struct {
		name       string
		s          *Stmt
		args       args
		wantResult driver.Rows
		wantErr    bool
	}{
		{
			"Success newRows", &Stmt{dsn: tsNewRows.URL, query: "SELECT 1 AS a, 2 AS b, 3 AS c"}, args{},
			&Rows{
				dsn:     tsNewRows.URL,
				columns: []string{"a", "b", "c"},
				types:   []elasticsearchType{ElaTypeInteger, ElaTypeInteger, ElaTypeInteger},
				rows:    [][]driver.Value{[]driver.Value{1, 2, 3}},
				cursor:  "",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.s.Query(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Stmt.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Stmt.Query() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
