package base58

import (
	"math/big"
)

const table = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var base = big.NewInt(58)

// Encode encodes bytes to base58 string
func Encode(src []byte) (res string) {
	x := new(big.Int).SetBytes(src)
	mod := new(big.Int)
	for x.Sign() > 0 {
		x.QuoRem(x, base, mod)
		res = string(table[mod.Int64()]) + res
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
