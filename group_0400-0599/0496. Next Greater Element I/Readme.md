## Problem

**Indonesia:**  
Elemen selanjutnya yang lebih besar dari suatu elemen x dalam sebuah array adalah elemen pertama yang lebih besar di sebelah kanan x pada array yang sama.

Diberikan dua array integer berbeda `nums1` dan `nums2` (0-indexed), di mana `nums1` adalah subset dari `nums2`.

Untuk setiap 0 <= i < nums1.length, temukan indeks j sehingga nums1[i] == nums2[j] dan tentukan elemen selanjutnya yang lebih besar dari nums2[j] di dalam nums2. Jika tidak ada elemen yang lebih besar, maka jawabannya adalah -1.

Kembalikan array `ans` dengan panjang nums1.length, di mana ans[i] adalah elemen selanjutnya yang lebih besar seperti dijelaskan di atas.

**English:**  
The next greater element of some element x in an array is the first greater element that is to the right of x in the same array.

You are given two distinct 0-indexed integer arrays `nums1` and `nums2`, where `nums1` is a subset of `nums2`.

For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the next greater element of nums2[j] in nums2. If there is no next greater element, then the answer for this query is -1.

Return an array `ans` of length nums1.length such that ans[i] is the next greater element as described above.

---

## Solution

**Indonesia:**  
Gunakan **monotonic decreasing stack** untuk memproses `nums2` dari kiri ke kanan.  
- Setiap kali menemukan elemen yang lebih besar dari elemen di atas stack, lakukan pop dan simpan pasangan (pop value, nilai sekarang) ke dalam hash table.
- Setelah selesai, untuk setiap elemen di `nums1`, jawabannya adalah hasil query dari hash table (jika tidak ada, jawab -1).

**Kompleksitas Waktu:** O(n + m), n = panjang nums2, m = panjang nums1  
**Kompleksitas Memori:** O(n)

**Contoh Implementasi (Go):**
```go
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    stack := []int{}
    hash := map[int]int{}

    for _, num := range nums2 {
        for len(stack) > 0 && num > stack[len(stack)-1] {
            hash[stack[len(stack)-1]] = num
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, num)
    }

    ans := make([]int, len(nums1))
    for i, num := range nums1 {
        if val, ok := hash[num]; ok {
            ans[i] = val
        } else {
            ans[i] = -1
        }
    }
    return ans
}
```

---

**English:**  
Use a **monotonic decreasing stack** to process `nums2` from left to right.  
- Whenever a greater element is found than the top of the stack, pop from the stack and store the pair (pop value, current value) in a hash table.
- After processing, for each element in `nums1`, the answer is the query result from the hash table (if not found, return -1).

**Time Complexity:** O(n + m), n = length of nums2, m = length of nums1  
**Space Complexity:** O(n)

**Sample Implementation (Go):**
```go
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    stack := []int{}
    hash := map[int]int{}

    for _, num := range nums2 {
        for len(stack) > 0 && num > stack[len(stack)-1] {
            hash[stack[len(stack)-1]] = num
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, num)
    }

    ans := make([]int, len(nums1))
    for i, num := range nums1 {
        if val, ok := hash[num]; ok {
            ans[i] = val
        } else {
            ans[i] = -1