func maxKDistinct(nums []int, k int) []int {
    sort.Slice(nums, func (a,b int) bool{
        return nums[a]>nums[b]
    })

    currIndex := 1
    res := []int{nums[0]}

    for i:= 1; i<len(nums)&&currIndex<k;i++{
        if nums[i] != nums[i-1] {
            res = append(res, nums[i])
            currIndex++
        }
    }

    return res
}