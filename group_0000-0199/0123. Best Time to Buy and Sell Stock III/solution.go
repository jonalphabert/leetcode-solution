func maxProfit(prices []int) int {
    buy1, sell1 := math.MinInt, 0
    buy2, sell2 := math.MinInt, 0
    
    for _, p := range prices {
        buy1 = max(buy1, -p)
        sell1 = max(sell1, buy1 + p)
        buy2 = max(buy2, sell1 - p)
        sell2 = max(sell2, buy2 + p)
    }

    return sell2
}