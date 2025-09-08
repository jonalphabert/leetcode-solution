func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }

    totalIncome := 0

    for i := 1; i<len(prices); i++ {
        if prices[i] > prices[i-1] {
            totalIncome += prices[i] - prices[i-1]
        }
    }

    return totalIncome
}