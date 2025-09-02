import (
    "sort"
)

func fourSum(nums []int, target int) [][]int {
    set := make(map[string]bool)
    result := make([][]int, 0)
    n := len(nums)

    sort.Slice(nums, func (l,r int) bool {
        return nums[l] < nums[r]
    })

    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            left := j+1
            right := n-1

            for left < right {
                currentRes := nums[i] + nums[j] + nums[left] + nums[right]

                if target == currentRes {
                    stringIndex := strconv.Itoa(nums[i]) + "," + strconv.Itoa(nums[j]) + "," + strconv.Itoa(nums[left]) + "," + strconv.Itoa(nums[right])

                    if !set[stringIndex] {
                        set[stringIndex] = true
                        result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
                    }
                    
                    left++
                    right--

                    for nums[left] == nums[left-1] {
                        left++
                        if left >= n {
                            break
                        }
                    }

                    for nums[right] == nums[right+1]{
                        right--
                        if right < 0 {
                            break
                        }
                    }
                } else if currentRes > target {
                    right --
                } else {
                    left ++
                }
            }
        }
    }

    return result
}