package base58

import (
	"encoding/hex"
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		want    string
		input   string
		wantErr bool
	}{
		{
			name:  "empty input",
			want:  "",
			input: "",
		},
		{
			want:  "61",
			input: "2g",
		},
		{
			want:  "626262",
			input: "a3gV",
		},
		{
			want:  "636363",
			input: "aPEr",
		},
		{
			want:  "73696d706c792061206c6f6e6720737472696e67",
			input: "2cFupjhnEsSn59qHXstmK2ffpLv2",
		},
		{
			want:  "00eb15231dfceb60925886b67d065299925915aeb172c06647",
			input: "1NS17iag9jJgTHD1VXjvLCEnZuQ3rJDE9L",
		},
		{
			want:  "516b6fcd0f",
			input: "ABnLTmg",
		},
		{
			want:  "572e4794",
			input: "3EFU7m",
		},
		{
			want:  "ecac89cad93923c02321",
			input: "EJDM8drfXA6uyA",
		},
		{
			want:  "10c8511e",
			input: "Rt5zm",
		},
		{
			want:  "00000000000000000000",
			input: "1111111111",
		},
		{
			want:  "801184cd2cdd640ca42cfc3a091c51d549b2f016d454b2774019c2b2d2e08529fd206ec97e",
			input: "5Hx15HFGyep2CfPxsJKe2fXJsCVn5DEiyoeGGF6JZjGbTRnqfiD",
		},
		{
			want:  "003c176e659bea0f29a3e9bf7880c112b1b31b4dc826268187",
			input: "16UjcYNBG9GTK4uq2f7yYEbuifqCzoLMGS",
		},
		{
			name:    "invalid input",
			input:   "-+",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "" {
				tt.name = tt.input
			}
			gotRes, err := Decode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if want := hex.EncodeToString(gotRes); want != tt.want {
				t.Errorf("Decode() = %v, want %v", want, tt.input)
			}
		})
	}
}
