func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
		return "0"
	}

    result := ""

    if (numerator < 0) != (denominator < 0) {
        result += "-"
    }

    num := abs(numerator)
    den := abs(denominator)

    result += strconv.Itoa(num/den)
    remainder := num % den

    if remainder == 0 {
        return result
    }

    result += "."

    seenRemainder := make(map[int]int)

    for remainder != 0 {
        if idx, ok := seenRemainder[remainder]; ok {
            return result[:idx] + "(" + result[idx:] + ")"
        }

        seenRemainder[remainder] = len(result)

        remainder *= 10
        result += strconv.Itoa(remainder/den)
        remainder %= den
    }

    return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}