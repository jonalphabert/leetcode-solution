## Problem

**Indonesia:**  
Terdapat n anak yang berdiri dalam satu barisan. Setiap anak memiliki nilai rating yang diberikan dalam array integer `ratings`.

Kamu harus membagikan permen kepada anak-anak tersebut dengan ketentuan:
- Setiap anak harus mendapat minimal satu permen.
- Anak dengan rating lebih tinggi harus mendapat lebih banyak permen daripada tetangganya.

Kembalikan jumlah minimum permen yang harus dibagikan agar semua aturan terpenuhi.

**English:**  
There are n children standing in a line. Each child is assigned a rating value given in the integer array `ratings`.

You are giving candies to these children subjected to the following requirements:
- Each child must have at least one candy.
- Children with a higher rating get more candies than their neighbors.

Return the minimum number of candies you need to have to distribute the candies to the children.

---

## Solution

### Bahasa Indonesia

Ide solusi:
1. Setiap anak diberi 1 permen terlebih dahulu.
2. Iterasi dari kiri ke kanan: jika rating anak di kanan lebih besar dari anak di kiri, maka jumlah permen anak kanan = jumlah permen anak kiri + 1.
3. Iterasi dari kanan ke kiri: jika rating anak di kiri lebih besar dari anak di kanan, maka jumlah permen anak kiri = maksimum antara jumlah permen anak kiri saat ini dan jumlah permen anak kanan + 1.
4. Jumlahkan semua permen yang telah diberikan.

Pendekatan ini memastikan aturan permen terpenuhi baik dari kiri maupun kanan.

### English

Solution idea:
1. Give each child 1 candy initially.
2. Iterate from left to right: if the rating of the right child is higher than the left, set the right child's candies to left child's candies + 1.
3. Iterate from right to left: if the rating of the left child is higher than the right, set the left child's candies to the maximum of its current candies and right child's candies + 1.
4. Sum all the candies distributed.

This approach ensures both left and right neighbor rules are satisfied.

**Time Complexity:** O(n)  
**Space Complexity:** O(n)