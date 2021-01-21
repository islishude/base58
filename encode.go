package base58

import (
	"math/big"
	"strings"
)

var base = big.NewInt(58)
var zero = new(big.Int)

const _table = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Encode encodes bytes to base58 string
func Encode(src []byte) (res string) {
	x := new(big.Int).SetBytes(src)
	mod := new(big.Int)
	for x.Cmp(zero) > 0 {
		x.QuoRem(x, base, mod)
		res = string(_table[mod.Int64()]) + res
	}

	for _, char := range src {
		if char == 0 {
			res = "1" + res
		} else {
			break
		}
	}
	return
}

func encodeChunk(raw []byte, padding int) (result string) {
	remainder := new(big.Int)
	remainder.SetBytes(raw)
	bigZero := new(big.Int)
	for remainder.Cmp(bigZero) > 0 {
		current := new(big.Int)
		remainder.DivMod(remainder, base, current)
		result = string(_table[current.Int64()]) + result
	}
	if len(result) < padding {
		result = strings.Repeat("1", (padding-len(result))) + result
	}
	return
}

func EncodeMonero(src []byte) (result string) {
	length := len(src)
	rounds := length / 8
	for i := 0; i < rounds; i++ {
		result += encodeChunk(src[i*8:(i+1)*8], 11)
	}
	if length%8 > 0 {
		result += encodeChunk(src[rounds*8:], 7)
	}
	return
}
