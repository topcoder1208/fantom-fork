package utils

import "math/big"

// ToZnx number of ZNX to Wei
func ToZnx(znx uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(znx), big.NewInt(1e18))
}
