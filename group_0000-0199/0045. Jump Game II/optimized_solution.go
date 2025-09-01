/* Information for this solution:
 * This solution inspired by Greedy algorithm which is more optimal than BFS.
 * The idea is to keep track of the farthest reachable index and the end of the current jump.
 * When we reach the end of the current jump, we increment the jump count and update the end to the farthest reachable index.
 * This proofs when we make solution on BFS way, the dp array always have the same value have pattern like 0,1,1,2,2,2,3,3,3,3...
 * So we can just count how many times we need to jump to reach the end of the array.
 * Complexity: O(n)
 * Space: O(1)
*/

func jump(nums []int) int {
    n := len(nums)
    farthest := 0
    jump := 0
    end := 0

    for i := 0; i < n-1; i++ {
        farthest = max(farthest, i+nums[i])
        if i == end {
            jump++
            end = farthest
        }
    }

    return jump
}