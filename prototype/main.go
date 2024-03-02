package main

import (
	"fmt"
	"math/big"
)

func main () {
	fmt.Println(out(300))
}

func out(n int) string {
	return dec(compute(n)).String()
}

func compute(n int) []byte {
	// n2 + n1 = n0
	r := [][]byte{
		{1, 0, 1},
		{1, 0, 1},
		{0, 1, 1},
	}
	for i := 1; i < n-1; i++ {
		rn := len(r) - 1
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
