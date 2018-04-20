package register

import (
	"reflect"
	"testing"
)

func TestRegister_ListRegistryConfig(t *testing.T) {
	tests := []struct {
		name string
		this *Register
		want []RegistryConf
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Register{}
			if got := this.ListRegistryConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register.ListRegistryConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister_Register(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		this *Register
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Register{}
			this.Register(tt.args.data)
		})
	}
}

func TestRegister_CreateNode(t *testing.T) {
	type args struct {
		registryConf RegistryConf
		data         string
	}
	tests := []struct {
		name string
		this *Register
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Register{}
			this.CreateNode(tt.args.registryConf, tt.args.data)
		})
	}
}

func TestRegister_UnRegister(t *testing.T) {
	tests := []struct {
		name string
		this *Register
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Register{}
			this.UnRegister()
		})
	}
}

func TestRegister_GetPath(t *testing.T) {
	tests := []struct {
		name string
		this *Register
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Register{}
			if got := this.GetPath(); got != tt.want {
				t.Errorf("Register.GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister_Close(t *testing.T) {
	tests := []struct {
		name string
		this *Register
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Register{}
			this.Close()
		})
	}
}
