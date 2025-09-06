func ceilDiv2(x int64) int64 {
    return (x + 1) / 2
}

// hitung F(n) = sum_{x=1..n} steps(x)
func F(n int64) int64 {
    if n <= 0 {
        return 0
    }

    var sum int64 = 0
    var start int64 = 1 // 4^(K-1)
    K := int64(1)

    // cari K sehingga start <= n < start*4
    for start*4 <= n {
        start *= 4
        K++
    }

    // sum blok penuh untuk i=1..K-1
    pow := int64(1)
    for i := int64(1); i <= K-1; i++ {
        blockCount := 3 * pow
        sum += i * blockCount
        pow *= 4
    }

    // sisa di blok K
    count := n - start + 1
    sum += K * count

    return sum
}


func minOperations(queries [][]int) int64 {
    var total int64 = 0
    for _, q := range queries {
        l := int64(q[0])
        r := int64(q[1])
        if r < l {
            l, r = r, l
        }

        diff := F(r) - F(l-1)
        ops := ceilDiv2(diff)
        total += ops
    }
    return total
}
