// Prefix Suffix Sum
func trap(height []int) int {
    length := len(height)
    arrMaxRight := make([]int, len(height))
    arrMaxLeft := make([]int, len(height))

    arrMaxLeft[0] = height[0]
    arrMaxRight[length-1] = height[length-1]

    for i := 1; i < length; i++ {
        arrMaxLeft[i] = max(arrMaxLeft[i-1], height[i])
    }

    for i := length-2; i >= 0; i-- {
        arrMaxRight[i] = max(arrMaxRight[i+1], height[i])
    }

    result := 0
    for i := 0; i < length; i++ {
        result += min(arrMaxRight[i], arrMaxLeft[i]) - height[i]
    }

    return result
}

// Monotomic Decreasing Stack
func trap(height []int) int {
    stack := []int{}
    result := 0

    for i := 0; i < len(height); i++ {
        for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
            top := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]

            if len(stack) == 0 {
                break
            }

            left := stack[len(stack) - 1]
            witdh := i - left - 1
            result += (min(height[left], height[i]) - height[top]) * witdh
        }

        stack = append(stack, i)
    }

    return result
}