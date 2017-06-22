package main

import (
	"fmt"
	"math/big"
)


// Compute the bezout coefficient u, v
// Condition : b < a
func extendedEuclidian(aa big.Int, bb big.Int) (big.Int, big.Int, error){
	var a, b, u0, v0, u1, v1, r, q, u, v, one big.Int
	a.Set(&aa)
	b.Set(&bb)
	
	u0.SetInt64(1)
	v0.SetInt64(0)
	u1.SetInt64(0)
	v1.SetInt64(1)
	one.SetInt64(1)
	
	if b.Cmp(&a) == 1 {
		return one, one, fmt.Errorf("a(%d) must be greater than b (%d)", a, b)
	}
	for {
		q.DivMod(&a,&b,&r)
		u.Mul(&q,&u1)
		u.Sub(&u0, &u)
		v.Mul(&q, &v1)
		v.Sub(&v0, &v)
		
		//shift
		u0.Set(&u1)
		v0.Set(&v1)
		u1.Set(&u)
		v1.Set(&v)
		
		if r.Cmp(&one) == 0 {
			break
		}
		a.Set(&b)
		b.Set(&r)
	}
	return u1, v1, nil
}
