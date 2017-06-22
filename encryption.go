package main

import (
	"fmt"
	"time"
	"math/big"
)


type Key interface {
	String() string
	Modulus() big.Int
	Exponent() big.Int
}

type SmallKey struct {
	modulus int64
	exponent int64
}

type BigKey struct {
	modulus big.Int
	exponent big.Int
}

func (k SmallKey) String() string {
	return fmt.Sprintf("Module: %d\nExponent: %d\n", k.modulus, k.exponent)
}

func (k SmallKey) Modulus() big.Int {
	var bm big.Int
	bm.SetInt64(k.modulus)
	return bm
}

func (k SmallKey) Exponent() big.Int {
	var be big.Int
	be.SetInt64(k.exponent)
	return be
}

func (k BigKey) String() string {
	return fmt.Sprintf("Module: %s\nExponent: %s\n", k.modulus.String(), k.exponent.String())
}

func (k BigKey) Modulus() big.Int {
	return k.modulus
}

func (k BigKey) Exponent() big.Int {
	return k.exponent
}

func generateKeys(strength int64) (Key, Key, error) {
	start := time.Now()
	defer func() { fmt.Printf("Keys generation last %s\n", time.Since(start))}()

	_, p := generateFermatPrime(1000)
	_, q := generateFermatPrime(1000)

	var bp, bq, bp1, bq1, bn, bfi, be, one big.Int
	one.SetInt64(1)
	bp.Set(&p)
	bq.Set(&q)
	bn.Mul(&bp, &bq)

	bp1.Sub(&bp, &one)
	bq1.Sub(&bq, &one)
	bfi.Mul(&bp1, &bq1)
	be.SetInt64(0)

	_, be = generateFermatPrime(1000)
	var check, prod big.Int
	_, bd, err := extendedEuclidian(bfi, be)
	if err != nil {
		return SmallKey{}, SmallKey{}, err
	}
	//if exponent is negative, build a positive one
	// d' = fi + d (as d is negative)
	if bd.Sign() == -1 {
		bd.Add(&bfi, &bd)
	}

	prod.Mul(&be, &bd)
	check.Mod(&prod, &bfi)

	return BigKey{bn, be}, BigKey{bn, bd}, nil
}


func encrypt(key Key, m int64) (big.Int, error) {
	start := time.Now()
	defer func() { fmt.Printf("Encryption last %s\n", time.Since(start))}()
	var bm, bn, be, crypto big.Int
	bm.SetInt64(m)
	bn = key.Modulus()
	be = key.Exponent()

	crypto.Exp(&bm, &be, &bn)
	return crypto, nil
}

func decrypt(key Key, bc big.Int) (big.Int, error) {
	start := time.Now()
	defer func() { fmt.Printf("Decryption last %s\n", time.Since(start))}()
	var  bn, bd, m big.Int
	bn = key.Modulus()
	bd = key.Exponent()

	m.Exp(&bc, &bd, &bn)
	return m, nil
}
