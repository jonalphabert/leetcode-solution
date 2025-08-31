## Problem

**Indonesia:**  
Diberikan sebuah deck kartu yang direpresentasikan sebagai array string `cards`, di mana setiap kartu berisi dua huruf kecil.  
Diberikan juga sebuah huruf `x`.  
Permainan dilakukan dengan aturan berikut:

- Mulai dengan 0 poin.
- Pada setiap giliran, cari dua kartu kompatibel dari deck yang **keduanya mengandung huruf x di posisi manapun**.
- Kartu kompatibel adalah dua kartu yang berbeda **tepat di satu posisi**.
- Hapus pasangan kartu tersebut dan dapatkan 1 poin.
- Permainan berakhir jika tidak ada pasangan kompatibel lagi.
- Kembalikan jumlah poin maksimum yang bisa didapat dengan strategi optimal.

**English:**  
You are given a deck of cards represented by a string array `cards`, and each card displays two lowercase letters.  
You are also given a letter `x`.  
You play a game with the following rules:

- Start with 0 points.
- On each turn, find two compatible cards from the deck that both contain the letter `x` in any position.
- Two cards are compatible if the strings differ in exactly 1 position.
- Remove the pair of cards and earn 1 point.
- The game ends when no more compatible pairs can be found.
- Return the maximum number of points you can gain with optimal play.

---

## Solution

### Bahasa Indonesia

1. **Klasifikasi Kartu:**
   - Kartu yang kedua hurufnya adalah `x` (`xx`): bisa dipasangkan dengan kartu lain yang mengandung `x` di satu posisi.
   - Kartu yang hanya huruf pertama adalah `x` (`x?`): simpan huruf kedua di map `cnt1`.
   - Kartu yang hanya huruf kedua adalah `x` (`?x`): simpan huruf pertama di map `cnt2`.

2. **Hitung Potensi Pasangan:**
   - Untuk kartu `x?`, pasangan kompatibel adalah dua kartu dengan huruf kedua berbeda.
   - Untuk kartu `?x`, pasangan kompatibel adalah dua kartu dengan huruf pertama berbeda.
   - Untuk kartu `xx`, bisa dipasangkan dengan kartu `x?` atau `?x` (karena hanya beda di satu posisi).

3. **Optimasi Pasangan:**
   - Untuk setiap kemungkinan jumlah kartu `xx` yang dipasangkan ke kiri (`x?`) dan ke kanan (`?x`), hitung maksimal pasangan yang bisa dibentuk.
   - Fungsi `solve(total, max)` menghitung maksimal pasangan dari grup, dengan menghindari pasangan yang hurufnya sama.
   - Coba semua pembagian kartu `xx` ke kiri dan kanan, ambil hasil maksimal.

4. **Implementasi:**
   - Hitung total dan maksimum frekuensi huruf di masing-masing grup.
   - Loop semua kemungkinan pembagian kartu `xx`, hitung skor maksimal.

### English

1. **Classify Cards:**
   - Cards with both letters as `x` (`xx`): can be paired with any card containing `x` in one position.
   - Cards with `x` as the first letter (`x?`): store the second letter in map `cnt1`.
   - Cards with `x` as the second letter (`?x`): store the first letter in map `cnt2`.

2. **Count Potential Pairs:**
   - For `x?` cards, compatible pairs are those with different second letters.
   - For `?x` cards, compatible pairs are those with different first letters.
   - For `xx` cards, they can be paired with `x?` or `?x` cards (since they differ in only one position).

3. **Optimize Pairing:**
   - For each possible distribution of `xx` cards to the left (`x?`) and right (`?x`), calculate the maximum pairs.
   - The function `solve(total, max)` computes the maximum pairs in a group, avoiding pairs with the same letter.
   - Try all possible splits of `xx` cards, and take the maximum score.

4. **Implementation:**
   - Count total and maximum frequency of letters in each group.
   - Loop through all possible splits of `xx` cards, and compute the maximal score.

**Time Complexity:** O(N)  
**Space Complexity:** O(N)