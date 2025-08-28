func lenOfVDiagonal(grid [][]int) int {
    dirs := [4][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
    m, n := len(grid), len(grid[0])
    memo := make([][][4][2]int, m)
    for i := range memo {
        memo[i] = make([][4][2]int, n)
        for j := range memo[i] {
            for k := range memo[i][j] {
                for l := range memo[i][j][k] {
                    memo[i][j][k][l] = -1
                }
            }
        }
    }

    var dfs func(cx, cy, direction int, turn int, target int) int
    dfs = func(cx, cy, direction int, turn int, target int) int {
        newX := cx + dirs[direction][0]
        newY := cy + dirs[direction][1]

        if newX < 0 || newX >= m || newY < 0 || newY >= n {
            return 0
        }

        if grid[newX][newY] != target {
            return 0
        }

        if memo[newX][newY][direction][turn] != -1 {
            return memo[newX][newY][direction][turn]
        }

        newTarget := 2
        if target == 2 {
            newTarget = 0
        }

        maxStep := dfs(newX, newY, direction, turn, newTarget)
        if turn == 0 {
            maxStep = max(maxStep, dfs(newX, newY, (direction+1)%4, 1, newTarget))
        }
        memo[newX][newY][direction][turn] = maxStep +1
        return maxStep + 1
    }

    res := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                for direction := 0; direction < 4; direction++ {
                    res = max(res, dfs(i, j, direction, 0, 2)+1)
                }
            }
        }
    }
    return res
}