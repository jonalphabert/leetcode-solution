func compareVersion(version1, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	
	maxLen := len(v1)
	if len(v2) > maxLen {
		maxLen = len(v2)
	}

	for i := 0; i < maxLen; i++ {
		r1 := getRevision(v1, i)
		r2 := getRevision(v2, i)

		if r1 > r2 {
			return 1
		}
		if r1 < r2 {
			return -1
		}
	}

	return 0
}

func getRevision(parts []string, idx int) int {
	if idx >= len(parts) {
		return 0
	}
	val, err := strconv.Atoi(parts[idx])
	if err != nil {
		return 0
	}
	return val
}