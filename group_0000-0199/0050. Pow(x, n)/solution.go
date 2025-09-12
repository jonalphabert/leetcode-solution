func fastPow(a float64, b int) float64 {
	res := float64(1)
	base := a
	exp := b

	for exp > 0 {
		if exp&1 == 1 { 
			res *= base
		}
		base *= base   
		exp >>= 1      
	}
	return res
}

func myPow(x float64, n int) float64 {
    if n < 0 {
        return float64(1)/fastPow(x, -1*n)
    }
    return fastPow(x, n)
}