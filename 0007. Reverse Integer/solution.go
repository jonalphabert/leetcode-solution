func reverse(x int) int {
    result := 0

    for x != 0 {
        if result > (1<<31-1)/10  || result < -(1<<31-1)/10 {
            return 0
        }
        result *= 10
        result += x%10
        x/=10
    }

    return result
}