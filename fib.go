package main

import (
	"math/big"
)

func NewFibo() Fibo {
	return fibo{}
}

type fibo struct{}

func (f fibo) ComputeAll(nth int) string {
	seq := f.compute(nth)
	// TODO: factors
	return dec(seq[len(seq)-1]).String()
}

func (f fibo) ComputeOnly(nth int) string {
	seq := f.computeOnly(nth)
	return dec(seq).String()
}

func (f fibo) compute(n int) [][]byte {
	// n2 + n1 = n0
	r := [][]byte{
		{1, 0, 1},
		{1, 0, 1},
		{0, 1, 1},
	}
	var sequences [][]byte
	for i := 1; i < n-1; i++ {
		rn := len(r) - 1
		// 20^0 r[2], 20^1 r[3], 20^2 r[3] ...
		powerUp := false
		for j := 0; j <= rn; j++ {
			// shit values to the left
			r[j][0] = r[j][1]
			r[j][1] = r[j][2]

			// n2 + n1 = n0
			c := r[j][0] + r[j][1]
			if powerUp {
				c++
				powerUp = false
			}
			if c >= 20 {
				c -= 20
				if j >= 2 {
					// if next row does not exist, create it
					if (j + 1) > rn {
						r = append(r, []byte{0, 0, 0})
						rn++
					}
					powerUp = true
				}
			}
			r[j][2] = c
		}
		// how can I flush the last the excess sequences and keep only the last 4?
		var seq []byte
		for i := 2; i <= len(r)-1; i++ {
			seq = append(seq, r[i][2])
		}
		sequences = append(sequences, seq)
	}
	return sequences}

func (f fibo) computeOnly(n int) []byte {
	// n2 + n1 = n0
	r := [][]byte{
		{1, 0, 1},
		{1, 0, 1},
		{0, 1, 1},
	}
	for i := 1; i < n-1; i++ {
		rn := len(r) - 1
		// 20^0 r[2], 20^1 r[3], 20^2 r[3] ...
		powerUp := false
		for j := 0; j <= rn; j++ {
			// shift values to the left
			r[j][0] = r[j][1]
			r[j][1] = r[j][2]

			// n2 + n1 = n0
			c := r[j][0] + r[j][1]
			if powerUp {
				c++
				powerUp = false
			}
			if c >= 20 {
				c -= 20
				if j >= 2 {
					// if next row does not exist, create it
					if (j + 1) > rn {
						r = append(r, []byte{0, 0, 0})
						rn++
					}
					powerUp = true
				}
			}
			r[j][2] = c
		}
	}
	var seq []byte
	for i := 2; i <= len(r)-1; i++ {
		seq = append(seq, r[i][2])
	}
	return seq
}

func dec(n []byte) *big.Int {
	power := len(n) - 1
	res := new(big.Int)
	for i := 0; i <= power; i++ {
		z := new(big.Int).Exp(big.NewInt(20), big.NewInt(int64(i)), nil)
		z = z.Mul(z, big.NewInt(int64(n[i])))
		res.Add(res, z)
	}
	return res
}
