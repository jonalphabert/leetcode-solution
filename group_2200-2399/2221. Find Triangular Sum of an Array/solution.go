// Brute force approach
// Time Complexity: O(n^2)
// Space Complexity: O(1)

// There is any O(n) approach using combinatorics and modular arithmetic and binomial coefficients
// I'll add that later

func triangularSum(nums []int) int {
    n := len(nums)
    for i:=0; i<n; i++ {
        for j:=0; j<n-i-1; j++ {
            nums[j] = (nums[j]+nums[j+1])%10
        }
    }

    return nums[0]
}