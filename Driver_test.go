package elasticsearch

import (
	"database/sql/driver"
	"reflect"
	"testing"
)

func TestDriver_Open(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		d       *Driver
		args    args
		want    driver.Conn
		wantErr bool
	}{
		{"Success", &Driver{}, args{"something"}, &Conn{"something"}, false},
		{"Success", &Driver{}, args{"http://localhost:9200"}, &Conn{"http://localhost:9200"}, false},
		{"Fail: Invalid argument value", &Driver{}, args{""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Open(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Driver.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Driver.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}
