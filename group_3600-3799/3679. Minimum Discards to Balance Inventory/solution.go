func minArrivalsToDiscard(arrivals []int, w int, m int) int {
    type itemInfo struct {
        val int
		day int
    }

	queueItem := []itemInfo{}
	itemValFreq := map[int]int{}
	discarded := 0

	for idx, arrival := range arrivals {
		// Drop barang yang sudah lewat hari pengiriman
		if len(queueItem) != 0 && queueItem[0].day <= idx - w {
			arriveItem := queueItem[0]
			queueItem = queueItem[1:]
			itemValFreq[arriveItem.val]--
		}

		// Cek frequency value apakah melebihi batas yang ditentukan atau tidak
		if itemValFreq[arrival] < m {
			queueItem = append(queueItem, itemInfo{val: arrival, day: idx})
			itemValFreq[arrival]++
		} else {
			discarded++
		}
	}

	return discarded
}