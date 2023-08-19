package model_test

import (
	"reflect"
	"testing"

	"github.com/aqyuki/initapp/model"
)

func TestNewConfig(t *testing.T) {
	var unexpected model.FileType = -1

	type args struct {
		ft   model.FileType
		path string
	}
	tests := []struct {
		name      string
		args      args
		want      *model.Config
		wantPanic bool
		panicMsg  string
	}{
		{
			name: "Normal JSON",
			args: args{
				ft:   model.FileTypeJSON,
				path: createTestPath(t),
			},
			want: &model.Config{
				Type: model.FileTypeJSON,
				Path: createTestPath(t),
			},
			wantPanic: false,
			panicMsg:  "",
		},
		{
			name: "Normal YAML",
			args: args{
				ft:   model.FileTypeYAML,
				path: createTestPath(t),
			},
			want: &model.Config{
				Type: model.FileTypeYAML,
				Path: createTestPath(t),
			},
			wantPanic: false,
			panicMsg:  "",
		},
		{
			name: "Normal TOML",
			args: args{
				ft:   model.FileTypeTOML,
				path: createTestPath(t),
			},
			want: &model.Config{
				Type: model.FileTypeTOML,
				Path: createTestPath(t),
			},
			wantPanic: false,
			panicMsg:  "",
		},
		{
			name: "Normal INI",
			args: args{
				ft:   model.FileTypeINI,
				path: createTestPath(t),
			},
			want: &model.Config{
				Type: model.FileTypeINI,
				Path: createTestPath(t),
			},
			wantPanic: false,
			panicMsg:  "",
		},
		{
			name: "Panic - Unexpected file format ID",
			args: args{
				ft:   unexpected,
				path: createTestPath(t),
			},
			want: &model.Config{
				Type: model.FileTypeINI,
				Path: createTestPath(t),
			},
			wantPanic: true,
			panicMsg:  "Unexpected file format ID specified",
		},
		{
			name: "Panic - Non file path",
			args: args{
				ft:   model.FileTypeJSON,
				path: "",
			},
			want: &model.Config{
				Type: model.FileTypeJSON,
				Path: "",
			},
			wantPanic: true,
			panicMsg:  "Path not specified",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err == nil {
					return
				}
				if !tt.wantPanic {
					t.Errorf("Error (%s) : Panic ensues., got -> %+v", tt.name, err)
				}
				if err != tt.panicMsg {
					t.Errorf("Error (%s) : PANIC message is different from expected, got -> %+v", tt.name, err)
				}
			}()
			got := model.NewConfig(tt.args.ft, tt.args.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error (%s) : want -> %+v , got -> %+v", tt.name, tt.want, got)
			}
		})
	}
}
