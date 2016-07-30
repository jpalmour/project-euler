package divisors

import (
	"fmt"
	"testing"
)

func TestGetDivisorCount(t *testing.T) {
	upperBound := 100
	dc := NewCalc(upperBound)
	for _, test := range []struct {
		num            int
		expectedCount  int
		expectedErrMsg string
	}{
		{5, 2, ""},
		{6, 4, ""},
		{12, 6, ""},
		{200, 0, "200 too large for divisors.Calc size 100"},
	} {
		count, err := dc.DivisorCount(test.num)
		if count != test.expectedCount {
			t.Errorf("dc.GetDivisorCount(%v), UpperBound=%v returned count of %v but expected %v", test.num, upperBound, count, test.expectedCount)
		}
		if err != nil && err.Error() != test.expectedErrMsg {
			t.Errorf("dc.GetDivisorCount(%v), UpperBound=%v expected error message: %v, but received error message: %v", test.num, upperBound, test.expectedErrMsg, err.Error())
		}
	}
}

func TestGetDivisors(t *testing.T) {
	upperBound := 100
	dc := NewCalc(upperBound)
	for _, test := range []struct {
		num         int
		expDivisors []int
		expErrMsg   string
	}{
		{5, []int{1, 5}, ""},
		{6, []int{1, 2, 3, 6}, ""},
		{12, []int{1, 2, 3, 4, 6, 12}, ""},
		{200, nil, "200 too large for divisors.Calc size 100"},
	} {
		divisors, err := dc.Divisors(test.num)
		if fmt.Sprintf("%v", divisors) != fmt.Sprintf("%v", test.expDivisors) {
			t.Errorf("dc.GetDivisors(%v)=%v, expected %v (UpperBound=%v)", test.num, divisors, test.expDivisors, upperBound)
		}
		if err != nil && err.Error() != test.expErrMsg {
			t.Errorf("dc.GetDivisors(%v), UpperBound=%v expected error message: %v, but received error message: %v", test.num, upperBound, test.expErrMsg, err.Error())
		}
	}
}

func benchmarkDivisorsBrute(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		dc := NewCalc(i)
		dc.DivisorsBrute(i)
	}
}

func benchmarkDivisors(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		dc := NewCalc(i)
		dc.Divisors(i)
	}
}

func BenchmarkDivisorsBrute100(b *testing.B) { benchmarkDivisorsBrute(100, b) }
func BenchmarkDivisors100(b *testing.B)      { benchmarkDivisors(100, b) }

func BenchmarkDivisorsBrute200(b *testing.B) { benchmarkDivisorsBrute(200, b) }
func BenchmarkDivisors200(b *testing.B)      { benchmarkDivisors(200, b) }

func BenchmarkDivisorsBrute405(b *testing.B) { benchmarkDivisorsBrute(405, b) }
func BenchmarkDivisors405(b *testing.B)      { benchmarkDivisors(405, b) }

func BenchmarkDivisorsBrute807(b *testing.B) { benchmarkDivisorsBrute(807, b) }
func BenchmarkDivisors807(b *testing.B)      { benchmarkDivisors(807, b) }

func BenchmarkDivisorsBrute1643(b *testing.B) { benchmarkDivisorsBrute(1643, b) }
func BenchmarkDivisors1643(b *testing.B)      { benchmarkDivisors(1643, b) }

func BenchmarkDivisorsBrute3284(b *testing.B) { benchmarkDivisorsBrute(3284, b) }
func BenchmarkDivisors3284(b *testing.B)      { benchmarkDivisors(3284, b) }

func BenchmarkDivisorsBrute6531(b *testing.B) { benchmarkDivisorsBrute(6531, b) }
func BenchmarkDivisors6531(b *testing.B)      { benchmarkDivisors(6531, b) }

func BenchmarkDivisorsBrute12977(b *testing.B) { benchmarkDivisorsBrute(12977, b) }
func BenchmarkDivisors12977(b *testing.B)      { benchmarkDivisors(12977, b) }

func BenchmarkDivisorsBrute25843(b *testing.B) { benchmarkDivisorsBrute(25843, b) }
func BenchmarkDivisors25843(b *testing.B)      { benchmarkDivisors(25843, b) }

func BenchmarkDivisorsBrute51028(b *testing.B) { benchmarkDivisorsBrute(51028, b) }
func BenchmarkDivisors51028(b *testing.B)      { benchmarkDivisors(51028, b) }

func BenchmarkDivisorsBrute102320(b *testing.B) { benchmarkDivisorsBrute(102320, b) }
func BenchmarkDivisors102320(b *testing.B)      { benchmarkDivisors(102320, b) }
