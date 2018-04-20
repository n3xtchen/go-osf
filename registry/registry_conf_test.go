package register

import (
	"reflect"
	"testing"
)

func TestNewRegistryConf(t *testing.T) {
	type args struct {
		id      string
		address string
	}
	tests := []struct {
		name string
		args args
		want *RegistryConf
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegistryConf(tt.args.id, tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegistryConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
