func sortMatrix(grid [][]int) [][]int {
    n := len(grid)
	diag := make(map[int][]int)

	// 1. Group by diagonal
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			key := i - j
			diag[key] = append(diag[key], grid[i][j])
		}
	}

	// 2. Sort by each diagonal
	for key, arr := range diag {
        // Karena di golang kita pengambilan hanya bisa dari index terakhir maka
        // diagonal bawah diurutkan asc supaya ketika diambil dari belakang membentuk urutan desc
        // sebaliknya diagonal atas diurutkan desc supaya ketika diambil dari belakang membentuk urutan asc
		if key >= 0 { 
			sort.Ints(arr)
		} else { 
            sort.Sort(sort.Reverse(sort.IntSlice(arr)))
		}
		diag[key] = arr
	}

	// 3. Reconstruct matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			key := i - j
			arr := diag[key]

			// Ambil elemen terakhir lalu hapus dari array
			val := arr[len(arr)-1]
			grid[i][j] = val
			diag[key] = arr[:len(arr)-1]
		}
	}

	return grid
}