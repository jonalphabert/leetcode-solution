const MOD = 1000000007

func peopleAwareOfSecret(n int, delay int, forget int) int {
    new := make([]int, n+1)
    new[1] = 1

    shareable := 0 

    for day := 2; day <= n; day++ {
        if day > delay {
            shareable = (shareable + new[day-delay]) % MOD
        }
        if day > forget {
            shareable = (shareable - new[day-forget] + MOD) % MOD
        }
        new[day] = shareable
    }

    ans := 0

    start := n-forget+1

    if start < 1 {
        start = 1
    }

    for day := n-forget+1; day <= n; day++ {
        ans = (ans + new[day]) % MOD
    }

    return ans
}