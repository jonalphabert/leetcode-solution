func score(cards []string, x byte) int {
    cnt1 := make(map[byte]int) 
    cnt2 := make(map[byte]int) 
    maxLeft := 0
    maxRight := 0
    totalLeft := 0
    totalRight := 0
    both := 0  

    for _, card := range cards {
        if card[0] == x && card[1] == x {
            both++
        } else if card[0] == x {
            cnt1[card[1]]++
        } else if card[1] == x {
            cnt2[card[0]]++
        }
    }

	for _, v := range cnt1 {
		totalLeft += v
		if v > maxLeft {
			maxLeft = v
		}
	}

	for _, v := range cnt2 {
		totalRight += v
		if v > maxRight {
			maxRight = v
		}
	}

    solve := func(total int, max int) int {
	return min(total/2, total-max)
	}

    // My missing piece, no looping when contest T.T
    maxPairs := 0
    for i := 0; i <= both; i++ {
        maxPairs = max(maxPairs, solve(totalLeft+i, max(maxLeft, i)) + solve(totalRight+both-i, max(maxRight, both-i)))
    }

    return maxPairs
}