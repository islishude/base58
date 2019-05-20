package base58

import "math/big"

const table = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var base = big.NewInt(58)
var zero = new(big.Int)

// Encode encodes bytes to base58 string
func Encode(src []byte) (res string) {
	x := new(big.Int).SetBytes(src)
	mod := new(big.Int)
	for x.Cmp(zero) > 0 {
		x.QuoRem(x, base, mod)
		res = string(table[mod.Int64()]) + res
	}

	for i := 0; i < len(src); i += 2 {
		if 0 == src[i] && 0 == src[i+1] {
			res = "1" + res
		} else {
			break
		}
	}
	return
}
