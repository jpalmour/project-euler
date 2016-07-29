package combinatorics

import (
	"math/big"
)

func Factorial(num *big.Int) *big.Int {
	if num.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(1)
	}
	temp := new(big.Int)
	temp.Add(num, big.NewInt(-1))
	temp = Factorial(temp)
	return temp.Mul(temp, num)
}
