package wubi

import (
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{name: "李启爱", args: "李启爱", want: []string{"sbf", "ynkd", "epdc"}},
		{name: "朱文倩", args: "朱文倩", want: []string{"rii", "yygy", "wgeg"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Get(tt.args); !StringSliceEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestGetChar(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{name: "", args: "启", want: "ynkd"},
		{name: "", args: "爱", want: "epdc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChar(tt.args)
			if err != nil {
				t.Errorf("GetChar() error = %v, wantErr %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("GetChar() got = %v, want %v", got, tt.want)
			}
		})
	}
}
