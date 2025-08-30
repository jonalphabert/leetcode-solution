func firstMissingPositive(nums []int) int {
    j := 0 

    // Partisi bilangan positif di kiri
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            tmp := nums[i]
            nums[i] = nums[j]
            nums[j] = tmp
            j++
        }
    }

    // Berusaha menempatkan bilangan positif pada tempatnya
	i := 0
    for i < j {
        x := nums[i]
		fmt.Println("if swap", nums, i, x)
        if x <= j && x != i+1 && nums[x-1] != x {
			fmt.Println("do swap", nums, i, x)
			fmt.Println("before swap", nums[i], nums[x-1], i, x)
            tmp := nums[i]
            nums[i] = nums[x-1]
            nums[x-1] = tmp
        } else {
            i++
        }
    }

	fmt.Println(nums)

    // Cek bilangan yang ada di pos, jika ditemukan berbeda maka return bilangan tersebut
    // Note : bilangan x akan ditempatkan di index x-1
    for i := 0; i < j; i++ {
        if nums[i] != i+1 {
            return i + 1
        }
    }
    return j + 1
}