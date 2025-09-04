const maxK = int(1e6 + 5)

func permutation(n, k int) int {
	if k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}

	res := int(1)
	for i := 1; i <= k; i++ {
		res = res * int(n-i+1) / int(i)
		// if res lebih dari maxK, return maxK
		if res >= maxK {
			return maxK
		}
	}
	return res
}

func countPalindrome(halfFreq []int) int {
	total := 0
	for _, f := range halfFreq {
		total += f
	}

	res := int(1)
	for _, cnt := range halfFreq {
		if cnt == 0 {
			continue
		}
		res *= permutation(total, cnt)
		// if res lebih dari maxK, return maxK
		// because k max 1e6 so the condition k >= maxK will never be true
		if res >= maxK {
			return maxK
		}
		total -= cnt
	}
	return res
}

func reverseString(s string) string {
    b := []byte(s)
    for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}

func smallestPalindrome(s string, k int) string {
    // buat frekuensi per charnya
    freq := make([]int, 26)

    for i := 0; i < len(s); i++ {
        freq[s[i]-'a']++
    }

    halfFreq := make([]int, 26)
    totalHalf := 0
    for i := 0; i < 26; i++ {
        halfFreq[i] = freq[i] / 2
        totalHalf += halfFreq[i]
    }

    if k > countPalindrome(halfFreq) {
        return ""
    }

    var halfBuilder strings.Builder
    for pos := 0; pos < totalHalf; pos++ {
        for i := 0; i < 26; i++ {
            if halfFreq[i] == 0 {
                continue
            }
            halfFreq[i]--
            cnt := countPalindrome(halfFreq)
            if k <= cnt {
                halfBuilder.WriteByte(byte('a' + i))
                break
            }
            k -= cnt
            halfFreq[i]++
        }
    }

    half := halfBuilder.String()
    mid := ""

    if len(s) % 2 == 1 {
        mid = string(s[len(s)/2])
    }

    return half + mid + reverseString(half)
}