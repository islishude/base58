package base58

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty input",
			input: "",
			want:  "",
		},
		{
			input: "61",
			want:  "2g",
		},
		{
			input: "626262",
			want:  "a3gV",
		},
		{
			input: "636363",
			want:  "aPEr",
		},
		{
			input: "73696d706c792061206c6f6e6720737472696e67",
			want:  "2cFupjhnEsSn59qHXstmK2ffpLv2",
		},
		{
			input: "00eb15231dfceb60925886b67d065299925915aeb172c06647",
			want:  "1NS17iag9jJgTHD1VXjvLCEnZuQ3rJDE9L",
		},
		{
			input: "516b6fcd0f",
			want:  "ABnLTmg",
		},
		{
			input: "572e4794",
			want:  "3EFU7m",
		},
		{
			input: "ecac89cad93923c02321",
			want:  "EJDM8drfXA6uyA",
		},
		{
			input: "10c8511e",
			want:  "Rt5zm",
		},
		{
			input: "00000000000000000000",
			want:  "1111111111",
		},
		{
			input: "801184cd2cdd640ca42cfc3a091c51d549b2f016d454b2774019c2b2d2e08529fd206ec97e",
			want:  "5Hx15HFGyep2CfPxsJKe2fXJsCVn5DEiyoeGGF6JZjGbTRnqfiD",
		},
		{
			input: "003c176e659bea0f29a3e9bf7880c112b1b31b4dc826268187",
			want:  "16UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGS",
		},
	}
	for _, tt := range tests {
		if tt.name == "" {
			tt.name = tt.input
		}

		t.Run(tt.name, func(t *testing.T) {
			input, err := hex.DecodeString(tt.input)
			if err != nil {
				t.Error(err)
				return
			}

			var zeroCount int

			for i := 0; i < len(tt.input); i += 2 {
				if tt.input[i] == '0' && tt.input[i+1] == '0' {
					zeroCount++
				} else {
					break
				}
			}

			var src []byte
			if zeroCount > 0 {
				src = bytes.Repeat([]byte{0x00}, zeroCount)
				src = append(src, input...)
			} else {
				src = input
			}
			if want := Encode(src); want != tt.want {
				t.Errorf("Encode() = %v, want %v", want, tt.want)
			}
		})
	}
}
