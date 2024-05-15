package vlc

import (
	"reflect"
	"testing"
)

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bks  BinaryChunks
		want string
	}{
		{
			name: "test1",
			bks:  BinaryChunks{BinaryChunk("00101111"), BinaryChunk("10000000")},
			want: "0010111110000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bks.Join(); got != tt.want {
				t.Errorf("BinaryChunks.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBinChunks(t *testing.T) {

	tests := []struct {
		name string
		args []byte
		want BinaryChunks
	}{
		{
			name: "test1",
			args: []byte{20, 30, 60, 18},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinChunks(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
