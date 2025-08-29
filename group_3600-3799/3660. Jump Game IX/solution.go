func maxValue(nums []int) []int {
    n := len(nums)
	if n == 0 {
		return []int{}
	}

	// prefix maximum
	prefMax := make([]int, n)
	prefMax[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > prefMax[i-1] {
			prefMax[i] = nums[i]
		} else {
			prefMax[i] = prefMax[i-1]
		}
	}

	// suffix minimum
	suffMin := make([]int, n)
	suffMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffMin[i+1] {
			suffMin[i] = nums[i]
		} else {
			suffMin[i] = suffMin[i+1]
		}
	}

	/* cari cut (pemisah komponen)
	 * Cut ini adalah indeks i dimana prefMax[i] <= suffMin[i+1]
	 * Dengan cut ini, kita bisa memastikan bahwa semua elemen di kiri cut
	 * (yaitu di blok [0..i]) pasti lebih kecil atau sama dengan semua elemen di kanan cut
	 * (yaitu di blok [i+1..n-1]). Dengan demikian, kita bisa mengubah semua elemen di blok kiri
	 * menjadi nilai maksimum di blok tersebut tanpa melanggar aturan.
	 */
	cuts := []int{}
	for i := 0; i < n-1; i++ {
		if prefMax[i] <= suffMin[i+1] {
			cuts = append(cuts, i)
		}
	}

	// assign hasil
	ans := make([]int, n)
	start := 0
	cuts = append(cuts, n-1)
	for _, c := range cuts {
		// cari max di blok [start..c]
		maxVal := nums[start]
		for j := start; j <= c; j++ {
			if nums[j] > maxVal {
				maxVal = nums[j]
			}
		}
		for j := start; j <= c; j++ {
			ans[j] = maxVal
		}
		start = c + 1
	}

	return ans
}