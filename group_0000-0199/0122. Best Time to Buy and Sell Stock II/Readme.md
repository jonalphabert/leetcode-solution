## 🇮🇩 Bahasa Indonesia

### 📌 Deskripsi Masalah

Diberikan array integer `prices` dimana `prices[i]` adalah harga saham pada hari ke-`i`.

Pada setiap hari, Anda dapat memutuskan untuk membeli dan/atau menjual saham. Anda hanya dapat memegang paling banyak satu saham pada waktu tertentu. Namun, Anda dapat membeli kemudian langsung menjual pada hari yang sama.

Temukan dan kembalikan keuntungan maksimum yang dapat dicapai.

---

### 🔍 Abstraksi & Ide

Solusi ini menggunakan pendekatan **greedy** (serakah). Ide utamanya adalah:

1. Kita dapat membeli saham pada hari `i-1` dan menjual pada hari `i` jika harga pada hari `i` lebih tinggi dari hari `i-1`
2. Dengan menangkap setiap kenaikan harga kecil, kita dapat mengakumulasi keuntungan maksimum
3. Pendekatan ini efektif karena kita diizinkan melakukan multiple transactions

---

### 💡 Penyelesaian

Implementasi menggunakan iterasi sederhana:

```go
func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }

    totalIncome := 0

    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            totalIncome += prices[i] - prices[i-1]
        }
    }

    return totalIncome
}
```

**Cara kerja:**
- Iterasi melalui array harga mulai dari indeks 1
- Bandingkan harga hari ini dengan harga kemarin
- Jika harga naik, tambahkan selisihnya ke total keuntungan
- Kembalikan total keuntungan yang terkumpul

---

### ⏱️ Kompleksitas

* **Waktu:** `O(n)` - hanya satu iterasi melalui array
* **Ruang:** `O(1)` - hanya menggunakan variabel konstan

---

## 🇬🇧 English

### 📌 Problem Description

You are given an integer array `prices` where `prices[i]` is the price of a given stock on the `i-th` day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.

---

### 🔍 Abstraction & Idea

This solution uses a **greedy approach**. The key insights are:

1. We can buy stock on day `i-1` and sell on day `i` if the price on day `i` is higher than day `i-1`
2. By capturing every small price increase, we can accumulate maximum profit
3. This approach works because we're allowed to make multiple transactions

---

### 💡 Solution

Simple iterative implementation:

```go
func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }

    totalIncome := 0

    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            totalIncome += prices[i] - prices[i-1]
        }
    }

    return totalIncome
}
```

**How it works:**
- Iterate through the price array starting from index 1
- Compare today's price with yesterday's price
- If price increases, add the difference to total profit
- Return the accumulated total profit

---

### ⏱️ Complexity

* **Time:** `O(n)` - single pass through the array
* **Space:** `O(1)` - uses only constant variables

---