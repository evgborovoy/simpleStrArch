package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "test_1",
			args: "HelLo, everyONe!",
			want: "!hel!lo, every!o!ne!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.args); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "test_1",
			args: "ten !net",
			want: "10011011000011001000100001011001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.args); got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {

	tests := []struct {
		name string
		args string
		want []byte
	}{
		{
			name: "test1",
			args: "Is ted Name!",
			want: []byte{33, 43, 205, 46, 68, 24, 116, 128},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
