// Implementation from geekforgeeks
// Read more: https://www.geeksforgeeks.org/dsa/print-all-possible-permutations-of-an-array-vector-without-duplicates-using-backtracking/

func permute(nums []int) [][]int {
    permutation := [][]int{}

    var perm func(arr []int, index int)
    perm = func(arr []int, index int) {
        if index == len(arr) {
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            permutation = append(permutation, tmp)
        } else {
            for i := index; i < len(arr); i++ {
                temp := arr[index]
                arr[index] = arr[i]
                arr[i] = temp

                perm(arr, index+1)

                temp = arr[index]
                arr[index] = arr[i]
                arr[i] = temp
            } 
        }
    }

    perm(nums, 0)

    return permutation
}