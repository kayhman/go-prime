package main

import (
	"fmt"
	"time"
	"math/big"
	"math/rand"
)


func fermatTest(p big.Int) bool {
	witness := []int64{2,3,5,7}
	var one, a, p_1, mod big.Int
	isPrime := true
	one.SetInt64(1)
	p_1.Sub(&p, &one)
	for _, w := range witness {
		a.SetInt64(w)
		mod.Exp(&a, &p_1, &p)
		if mod.Cmp(&one) != 0 {
			isPrime = false
			break
		}
	}
	return isPrime
}


func generateFermatPrime(maxTry int64) (int64, big.Int) {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var p, min, max, diff, n, zero, one, two big.Int
	zero.SetInt64(0)
	one.SetInt64(1)
	two.SetInt64(2)
	n.SetInt64(1023)
	min.Exp(&two, &n, nil)
	n.SetInt64(1024)
	max.Exp(&two, &n, nil)
	max.Sub(&max, &one)
	diff.Sub(&max, &min)
	for i := int64(0) ; i < maxTry ; i++ {
		var p, even big.Int
		p.Rand(rd, &diff)
		p.Add(&p, &min)
		even.Mod(&p, &two)
		if even.Cmp(&zero) == 0 { // p is even
			p.Add(&p, &one)
		}
		if fermatTest(p) {
			//fmt.Printf("found prime in %d %s (%d) -> %v\n", i, p.String(), len(p.Bytes()) * 8, p.ProbablyPrime(10))
			return i, p
		}
	}
	return maxTry, p
}

func contains(primes []big.Int, e big.Int) bool {
	for _, p := range primes {
		if p.Cmp(&e) == 0 {
			return true
		}
	}
	return false
}

//0.003455
//Prime number density is 0.003400 : 9410 / %!d(float64=2.76756e+06)
// Prime number density is 0.003477 : 9399 / 2703294
//Prime number density is 0.003423 : 9400 / 2746224
//Prime number density is 0.003405 : 9437 / 2771571
//Prime number density is 0.003462 : 9429 / 2723856
//Prime number density is 0.003464 : 9373 / 2705767
//Prime number density is 0.003449 : 9379 / 2719328
func primeStat(maxTests int64) {
	count := int64(0)
	already := int64(0)
	primes := make([]big.Int, 0)
	for i := int64(0) ; i < maxTests ; i++ {
		nbTests, p := generateFermatPrime(1000)
		if nbTests == 1000 {
			fmt.Printf("Failed to found prime !\n")
			continue
		}
		if contains(primes, p) {
			fmt.Printf("prime already found: %d\n", already)
			already++
		} else {
			primes = append(primes,p)
			count += nbTests
		}
	}
	fmt.Printf("Prime number density is %f : %d / %d\n", float64(len(primes))/float64(count), len(primes),  count)
}
