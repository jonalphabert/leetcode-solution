func search(nums []int, target int) int {
    size := len(nums)
    left := 0
    right := size-1
    mid := 0

    // Mencari nilai terkecil dari array
    for left < right {
        mid = left + (right - left)/2

        if nums[left] < nums[right] {
            break
        }

        // Jika angka di mid masih lebih besar daripada angka di left
        // maka bisa dipastikan semua nilai di sebelah kiri mid sudah terurut
        // Jika kebalikannya, maka partisi dari array yang terrotasi ada di antara left dan juga mid
        if nums[left] <= nums[mid] {
            left = mid + 1
        } else {
            right = mid
        }
    }

    // Penentuan pivot left
    // Jika left adalah nol maka data telah terurut
    // Jika left tidak nol maka ada partisi
    if left != 0 {
        if nums[0] > target {
            right = size - 1
        } else {
            right = left-1
            left = 0
        }
    }

    // pencarian dari partisi yang ditemukan
    for left <= right {
        mid = left + (right - left) /2

        if nums[mid] == target {
            return mid
        }

        if target < nums[mid] {
            right = mid-1
        } else {
            left = mid+1
        }
    }

    return -1
}