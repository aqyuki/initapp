package model_test

import (
	"reflect"
	"testing"

	"github.com/aqyuki/initapp/model"
)

func TestNewConverter(t *testing.T) {

	type sample struct {
		Name string
		Num  int
	}

	createSampleStruct := func() *sample {
		return &sample{
			Name: "sample",
			Num:  1,
		}
	}

	type args struct {
		target interface{}
		config *model.Config
	}
	tests := []struct {
		name      string
		args      args
		want      *model.Converter
		wantPanic bool
		panicMsg  string
	}{
		{
			name: "Normal",
			args: args{
				target: createSampleStruct(),
				config: model.NewConfig(model.FileTypeJSON, createTestPath(t)),
			},
			want: &model.Converter{
				Target: createSampleStruct(),
				Config: model.NewConfig(model.FileTypeJSON, createTestPath(t)),
			},
			wantPanic: false,
			panicMsg:  "",
		},
		{
			name: "Panic - nil target",
			args: args{
				target: nil,
				config: model.NewConfig(model.FileTypeJSON, createTestPath(t)),
			},
			want: &model.Converter{
				Target: nil,
				Config: model.NewConfig(model.FileTypeJSON, createTestPath(t)),
			},
			wantPanic: true,
			panicMsg:  "target structure is nil",
		},
		{
			name: "Panic - nil config",
			args: args{
				target: createSampleStruct(),
				config: nil,
			},
			want: &model.Converter{
				Target: createSampleStruct(),
				Config: nil,
			},
			wantPanic: true,
			panicMsg:  "config is nil",
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
			got := model.NewConverter(tt.args.target, tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error (%s) : want -> %+v , got -> %+v", tt.name, tt.want, got)
			}
		})
	}
}
