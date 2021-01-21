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
		weight := strings.IndexByte(_table, char)
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

var _lookup = map[rune]int64{
	'1': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8, 'A': 9,
	'B': 10, 'C': 11, 'D': 12, 'E': 13, 'F': 14, 'G': 15, 'H': 16, 'J': 17, 'K': 18,
	'L': 19, 'M': 20, 'N': 21, 'P': 22, 'Q': 23, 'R': 24, 'S': 25, 'T': 26, 'U': 27,
	'V': 28, 'W': 29, 'X': 30, 'Y': 31, 'Z': 32, 'a': 33, 'b': 34, 'c': 35, 'd': 36,
	'e': 37, 'f': 38, 'g': 39, 'h': 40, 'i': 41, 'j': 42, 'k': 43, 'm': 44, 'n': 45,
	'o': 46, 'p': 47, 'q': 48, 'r': 49, 's': 50, 't': 51, 'u': 52, 'v': 53, 'w': 54,
	'x': 55, 'y': 56, 'z': 57,
}

func decodeChunk(encoded string) []byte {
	bigResult := big.NewInt(0)
	currentMultiplier := big.NewInt(1)
	tmp := new(big.Int)

	for _, v := range encoded {
		index, valid := _lookup[v]
		if !valid {
			return nil // TODO: return error
		}

		tmp.SetInt64(index)
		tmp.Mul(currentMultiplier, tmp)
		bigResult.Add(bigResult, tmp)
		currentMultiplier.Mul(currentMultiplier, base)
	}

	return bigResult.Bytes()
}

func DecodeMonero(input string) (result []byte) {
	length := len(input)
	rounds := length / 11
	for i := 0; i < rounds; i++ {
		result = append(result, decodeChunk(input[i*11:(i+1)*11])...)
	}
	if length%11 > 0 {
		result = append(result, decodeChunk(input[rounds*11:])...)
	}
	return
}
