package algorithms

import "fmt"

func iSqrt(n int64) int64 {
	var r1, r int64 = n, n + 1
	for r1 < r {
		r, r1 = r1, (r1+n/r1)>>1
	}
	return r
}

func PrimeSieve(n int64) []int64 {
	result := make([]int64, 0, n)
	sieve := make([]bool, n+1)
	sn := iSqrt(n)
	var i, j int64
	for i = 2; i <= sn; i++ {
		if !sieve[i] {
			for j = i * i; j <= n; j += i {
				sieve[j] = true
			}
		}
	}
	for i = 2; i <= n; i++ {
		if !sieve[i] {
			result = append(result, i)
		}
	}
	return result
}

func EulersTotient(max uint64) []uint64 {
	phi := make([]uint64, max+1)
	phi[1] = 1
	for n := uint64(2); n <= max; n++ {
		// The "for i" loop below handles factors of N.
		// If phi(N) == 0 right here, then N is prime.
		if phi[n] == 0 {

			// phi(prime) == prime - 1
			phi[n] = n - 1

			// Iterate over 2N, 3N, 4N ... max
			for i := n << 1; i <= max; i += n {
				// If zero, initialize to itself
				if phi[i] == 0 {
					phi[i] = i
				}
				// Incrementally calculate phi(i)
				phi[i] = phi[i] / n * (n - 1)
			}
		}
	}
	return phi
}
func primes() {
	fmt.Println(EulersTotient(10))
	// primes := PrimeSieve(1000000000)
	// fmt.Println(len(primes))
}
