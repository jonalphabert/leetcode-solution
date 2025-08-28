func finalPrices(prices []int) []int {
    stack := []int{}

    for i:=0; i < len(prices) ; i++ {
        for len(stack) > 0 && prices[i] <= prices[stack[len(stack)-1]] {
            top := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]

            prices[top] -= prices[i]
        }

        stack = append(stack, i)
    }

    return prices
}