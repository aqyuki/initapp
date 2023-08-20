package initapp

import (
	"reflect"
	"testing"
)

func Test_decodeJSON(t *testing.T) {
	type test struct {
		A string
		B string
	}
	type args struct {
		raw []byte
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Normal",
			args: args{
				raw: []byte(`{"A":"sample","B":"hello"}`),
			},
			want: test{
				A: "sample",
				B: "hello",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeJSON[test](tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("(%s) error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("%s : %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
