func isValidSudoku(board [][]byte) bool {
    rows := make([]map[byte]bool, 9)
    cols := make([]map[byte]bool, 9)
    grids := make([]map[byte]bool, 9)

    for i := range 9 {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        grids[i] = make(map[byte]bool)
    }

    for i := 0; i < 9; i++ {
        for j := 0; j<9; j++ {
            val := board[i][j]

            if val == '.' {
                continue
            }

            gridIdx := (i/3) * 3 + j/3
            // validasi kalau kosong mapnya
            if rows[i][val] == true || cols[j][val] == true || grids[gridIdx][val] == true {
                return false
            }

            rows[i][val] = true
            cols[j][val] = true
            grids[gridIdx][val] = true
        }
    }

    return true
}