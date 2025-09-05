func fastPow(a, b int) int {
	res := 1
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

func countBit(num int) int {
    count := 0

    for num > 0 {
		// fmt.Println("num:", num, " count:", count)
        count += num % 2
        num >>= 1
    }

    return count
}

func makeTheIntegerZero(num1 int, num2 int) int {
    k := 1
    rem := num1 - (k * num2)
    maxRem := fastPow(2, 61)

    for rem > 0 && rem < maxRem {
		// fmt.Println("k:", k, " rem:", rem, " countBit(rem):", countBit(rem))
        if countBit(rem) <= k && k <= rem {
			// fmt.Println("Final k:", k)
            return k
        }
        k++
        rem = num1 - (k * num2)
    }

	// fmt.Println("Exited loop with k:", k, " rem:", rem, " countBit(rem):", countBit(rem))
    return -1
}