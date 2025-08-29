func flowerGame(n int, m int) int64 {
    oddN := (n + 1) / 2   
    evenN := n / 2        
    oddM := (m + 1) / 2   
    evenM := m / 2        

    ans := int64(oddN*evenM + evenN*oddM)
    return ans
}