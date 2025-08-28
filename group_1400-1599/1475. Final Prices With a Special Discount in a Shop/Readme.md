# Indonesia
## Problem

Diberikan sebuah array integer `prices` di mana `prices[i]` adalah harga barang ke-i di toko.

Terdapat diskon khusus untuk barang di toko. Jika kamu membeli barang ke-i, maka kamu akan mendapatkan diskon sebesar `prices[j]` di mana `j` adalah indeks minimum sehingga `j > i` dan `prices[j] <= prices[i]`. Jika tidak ada indeks seperti itu, maka kamu tidak mendapatkan diskon.

Kembalikan array integer `answer` di mana `answer[i]` adalah harga akhir yang harus dibayar untuk barang ke-i, dengan mempertimbangkan diskon khusus.

**Contoh:**
```
Input: prices = [8,4,6,2,3]
Output: [4,2,4,2,3]
Penjelasan:
- Barang pertama: diskon dari barang ke-2 (4), jadi 8-4=4
- Barang kedua: diskon dari barang ke-4 (2), jadi 4-2=2
- Barang ketiga: diskon dari barang ke-4 (2), jadi 6-2=4
- Barang keempat: tidak ada diskon, tetap 2
- Barang kelima: tidak ada diskon, tetap 3
```

## Solution

Solusi menggunakan **Monotonic Stack**:

- Iterasi setiap harga barang.
- Gunakan stack untuk menyimpan indeks barang yang belum mendapatkan diskon.
- Jika harga saat ini lebih kecil atau sama dengan harga pada indeks teratas stack, maka barang pada indeks tersebut mendapatkan diskon dari harga saat ini.
- Update harga barang pada indeks tersebut dan keluarkan dari stack.
- Tambahkan indeks barang saat ini ke stack.
- Kembalikan array harga yang sudah diperbarui.

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Memori:** O(n)

**Contoh Implementasi (Go):**
```go
func finalPrices(prices []int) []int {
    stack := []int{}

    for i := 0; i < len(prices); i++ {
        for len(stack) > 0 && prices[i] <= prices[stack[len(stack)-1]] {
            top := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]

            prices[top] -= prices[i]
        }

        stack = append(stack, i)
    }

    return prices
}
```

---

# English
## Problem

You are given an integer array `prices` where `prices[i]` is the price of the ith item in a shop.

There is a special discount for items in the shop. If you buy the ith item, you will receive a discount equal to `prices[j]` where `j` is the minimum index such that `j > i` and `prices[j] <= prices[i]`. Otherwise, you will not receive any discount.

Return an integer array `answer` where `answer[i]` is the final price you will pay for the ith item of the shop, considering the special discount.

**Example:**
```
Input: prices = [8,4,6,2,3]
Output: [4,2,4,2,3]
Explanation:
- First item: discount from item 2 (4), so 8-4=4
- Second item: discount from item 4 (2), so 4-2=2
- Third item: discount from item 4 (2), so 6-2=4
- Fourth item: no discount, stays 2
- Fifth item: no discount, stays 3
```

## Solution

The solution uses a **Monotonic Stack**:

- Iterate through each item's price.
- Use a stack to store indices of items that have not received a discount yet.
- If the current price is less than or equal to the price at the top index of the stack, the item at that index receives a discount from the current price.
- Update the price for that item and remove it from the stack.
- Add the current item's index to the stack.
- Return the updated prices array.

**Time Complexity:** O(n)  
**Space Complexity:** O(n)

**Sample Implementation (Go):**
```go
func finalPrices(prices []int) []int {
    stack := []int{}

    for i := 0; i < len(prices); i++ {
        for len(stack) > 0 && prices[i] <= prices[stack[len(stack)-1]] {
            top := stack[len(stack) - 1]
            stack = stack[:len(stack)-1]

            prices[top] -= prices[i]
        }

        stack = append(stack, i)
    }
    return prices
}
```