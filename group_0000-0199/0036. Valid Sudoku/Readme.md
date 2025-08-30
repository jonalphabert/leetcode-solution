## Problem

**Indonesia:**  
Tentukan apakah sebuah papan Sudoku berukuran 9 x 9 valid.  
Hanya sel yang sudah terisi yang perlu divalidasi dengan aturan berikut:

- Setiap baris harus berisi digit 1-9 tanpa pengulangan.
- Setiap kolom harus berisi digit 1-9 tanpa pengulangan.
- Setiap kotak 3 x 3 di grid harus berisi digit 1-9 tanpa pengulangan.

Catatan:
- Papan Sudoku (yang terisi sebagian) bisa saja valid tetapi belum tentu dapat diselesaikan.
- Hanya sel yang sudah terisi yang perlu divalidasi sesuai aturan di atas.

**English:**  
Determine if a 9 x 9 Sudoku board is valid.  
Only the filled cells need to be validated according to the following rules:

- Each row must contain the digits 1-9 without repetition.
- Each column must contain the digits 1-9 without repetition.
- Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:
- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.

---

## Solution

### Bahasa Indonesia

- Buat 3 array map untuk menyimpan digit yang sudah muncul di setiap baris, kolom, dan kotak 3x3.
- Iterasi seluruh sel pada papan. Jika sel kosong ('.'), lanjutkan.
- Untuk setiap sel terisi, cek apakah digit sudah pernah muncul di baris, kolom, atau kotak 3x3 yang sesuai.
- Jika ada pengulangan, return `false`.
- Jika tidak ada pengulangan di seluruh papan, return `true`.

**Kompleksitas Waktu:** O(1) (karena ukuran papan tetap 9x9)  
**Kompleksitas Memori:** O(1)

### English

- Create 3 arrays of maps to store digits that have appeared in each row, column, and 3x3 box.
- Iterate through all cells on the board. If the cell is empty ('.'), continue.
- For each filled cell, check if the digit has already appeared in the corresponding row, column, or 3x3 box.
- If any repetition is found, return `false`.
- If no repetition is found for all filled cells, return `true`.

**Time Complexity:** O(1) (since the board size is fixed at 9x9)  
**Space Complexity:** O(1)

**Sample Implementation (Go):**
```go
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
        for j := 0; j < 9; j++ {
            val := board[i][j]
            if val == '.' {
                continue
            }
            gridIdx := (i/3)*3 + j/3
            if rows[i][val] || cols[j][val] || grids[gridIdx][val] {
                return false
            }
            rows[i][val] = true
            cols[j][val] = true
            grids[gridIdx][val] = true
        }
    }
    return true
}
```