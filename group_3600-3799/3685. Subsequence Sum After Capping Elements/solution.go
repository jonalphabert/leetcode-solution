func subsequenceSumAfterCapping(nums []int, sum int) []bool {
    n := len(nums)
    ans := make([]bool, len(nums))
    dp := make([]bool, sum+1)
    dp[0] = true
    i := 0

    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    for x:=1; x<=n; x++ {
        // Untuk semua angka yang dibawah x maka dp angka tersebut tidak akan berubah
        for i<n && nums[i] < x {
            for c:=sum; c>=nums[i]; c-- {
                dp[c] = dp[c-nums[i]] || dp[c]
            }
            i++
        }

        // Karena semua bilangan dibawah x sudah di precompute maka kita dapat mencari sum dari penyusunan x sebanyak k
        for k := 0; k <= n-i && sum-k*x >= 0; k++ {
            if dp[sum-k*x] {
                ans[x-1] = true
                break
            }
        }
    }
    
    return ans
}