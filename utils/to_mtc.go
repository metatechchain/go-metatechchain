package utils

import "math/big"

// ToMtc number of MTC to Wei
func ToMtc(mtc uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(mtc), big.NewInt(1e18))
}
