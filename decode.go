package base58

import (
	"errors"
	"math/big"
	"strings"
)

// Decode decoded base58 string to bytes
func Decode(input string) (res []byte, err error) {
	result := big.NewInt(0)
	isBreak := false
	for i := 0; i < len(input); i++ {
		char := input[i]
		weight := strings.IndexByte(table, char)
		if weight == -1 {
			return nil, errors.New("Invalid character in input string")
		}
		result.Mul(result, base)
		result.Add(result, big.NewInt(int64(weight)))

		if !isBreak {
			if i-1 > 0 && input[i-1] != '1' {
				isBreak = true
				continue
			}
			if char == '1' {
				res = append(res, 0)
			}
		}
	}
	res = append(res, result.Bytes()...)
	return
}
