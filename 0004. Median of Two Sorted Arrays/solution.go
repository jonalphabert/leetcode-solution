func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	n, m := len(nums1), len(nums2)
	l, r := 0, n
	half := (n + m) / 2
	total := n + m

	for l <= r {
		i := (l + r) / 2
		j := half - i

		var left1, right1, left2, right2 float64

		if i > 0 {
			left1 = float64(nums1[i-1])
		} else {
			left1 = math.Inf(-1)
		}

		if i < n {
			right1 = float64(nums1[i])
		} else {
			right1 = math.Inf(1)
		}

		if j > 0 {
			left2 = float64(nums2[j-1])
		} else {
			left2 = math.Inf(-1)
		}

		if j < m {
			right2 = float64(nums2[j])
		} else {
			right2 = math.Inf(1)
		}

		if left1 <= right2 && left2 <= right1 {
			if total%2 == 0 {
				return (max(left1, left2) + min(right1, right2)) / 2
			} else {
				return min(right1, right2)
			}
		}

		if left1 > right2 {
			r = i - 1
		} else {
			l = i + 1
		}
	}

	// fallback (harus ada supaya compile)
	return -1
}