package primes

var primesCache = []int{}

type primes struct {
	count, value int
}

func newPrimes() primes {
	return primes{count: 1, value: 2}
}

func (t *primes) next() {
	t.count++
	if t.count <= len(primesCache) {
		t.value = primesCache[t.count-1]
	} else {
		// TODO: replace with prime sieve
		for t.value++; true; t.value++ {
			if isPrime(t.value) {
				return
			}
		}
	}
	return
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func GetPrimeFactorsMap(num int) map[int]int {
	primeFactors := make(map[int]int)
	for p := newPrimes(); num > 1; p.next() {
		for num%p.value == 0 {
			num /= p.value
			primeFactors[p.value]++
		}
	}
	return primeFactors
}
