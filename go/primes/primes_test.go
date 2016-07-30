package primes

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	upperBound := 50
	pc := NewCalc(upperBound)
	for _, test := range []struct {
		num       int
		exp       bool
		expErrMsg string
	}{
		{5, true, ""},
		{6, false, ""},
		{12, false, ""},
		{13, true, ""},
		{20, false, ""},
		{41, true, ""},
		{200, false, "200 too large for primes.Calc size 50"},
	} {
		val, err := pc.IsPrime(test.num)
		if val != test.exp {
			t.Errorf("pc.IsPrime(%v), UpperBound=%v returned %v but expected %v", test.num, upperBound, val, test.exp)
		}
		if err != nil && err.Error() != test.expErrMsg {
			t.Errorf("dc.GetDivisorCount(%v), UpperBound=%v expected error message: %v, but received error message: %v", test.num, upperBound, test.expErrMsg, err.Error())
		}
	}
}

func TestNewCalc(t *testing.T) {
	primes := NewCalc(40).Primes
	expPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
	if fmt.Sprintf("%v", primes) != fmt.Sprintf("%v", expPrimes) {
		t.Errorf("NewCalc(40).Primes=%v, but expected %v", primes, expPrimes)
	}
}

func benchmarkPrimes(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewCalc(i)
	}
}

func BenchmarkPrimes100(b *testing.B)  { benchmarkPrimes(100, b) }
func BenchmarkPrimes200(b *testing.B)  { benchmarkPrimes(200, b) }
func BenchmarkPrimes400(b *testing.B)  { benchmarkPrimes(400, b) }
func BenchmarkPrimes800(b *testing.B)  { benchmarkPrimes(800, b) }
func BenchmarkPrimes1600(b *testing.B) { benchmarkPrimes(1600, b) }
func BenchmarkPrimes3200(b *testing.B) { benchmarkPrimes(3200, b) }
