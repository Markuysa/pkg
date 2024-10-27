package config

import (
	"reflect"
	"testing"
)

func TestLoadFromYAML(t *testing.T) {
	type config struct {
		Field1 string `yaml:"field1"`
		Field2 string `yaml:"field2"`
	}
	type args struct {
		cfg  *config
		path string
	}

	tests := []struct {
		name         string
		args         args
		expectConfig *config
		wantErr      bool
	}{
		{
			name: "SuccessLoadYAML",
			args: args{
				cfg:  &config{},
				path: "./test.yaml",
			},
			expectConfig: &config{
				Field1: "test",
				Field2: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := LoadFromYAML(tt.args.cfg, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFromYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.args.cfg, tt.expectConfig) {
				t.Errorf("LoadFromYAML() got = %v, want %v", tt.args.cfg, tt.expectConfig)
			}
		})
	}
}
