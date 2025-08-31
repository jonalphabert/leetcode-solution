func recoverOrder(order []int, friends []int) []int {
    posIndex := make(map[int]int, len(order))
	for i, id := range order {
		posIndex[id] = i
	}

	sort.Slice(friends, func(i, j int) bool {
		return posIndex[friends[i]] < posIndex[friends[j]]
	})

	return friends
}