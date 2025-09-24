func totalNQueens(n int) int {
    var res = 0
    if n==1 {
        return 1
    }
    if n==2 || n==3 {
        return 0
    }

    board := make([][]byte, n)

    for i := 0; i < n; i++ {
        board[i] = make([]byte, n)
        for j := 0; j < n; j++ {
            board[i][j] = '.'
        }
    }

    // Backtracking from here
    var trackQueen func(board [][]byte, row int)
    trackQueen = func(board [][]byte, row int) {
        if row == n {
            res++
            return
        }

        for col := 0; col < n; col++ {
            if isSafe(board, row, col, n) {
                board[row][col] = 'Q'
                trackQueen(board, row+1)
                board[row][col] = '.' // backtrack
            }
        }
    }

    trackQueen(board, 0)

    return res
}

func isSafe(board [][]byte, row, col, n int) bool {
    // cek kolom
    for i := 0; i < row; i++ {
        if board[i][col] == 'Q' {
            return false
        }
    }

    // cek diagonal kiri atas
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == 'Q' {
            return false
        }
    }

    // cek diagonal kanan atas
    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if board[i][j] == 'Q' {
            return false
        }
    }

    return true
}