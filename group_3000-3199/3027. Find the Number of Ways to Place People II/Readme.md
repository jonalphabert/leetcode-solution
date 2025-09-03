## Problem

**Indonesia:**  
Diberikan sebuah array 2D `points` berukuran n x 2 yang merepresentasikan koordinat integer dari beberapa titik pada bidang 2D, di mana `points[i] = [xi, yi]`.

Kamu harus menempatkan n orang (termasuk Alice dan Bob) di titik-titik tersebut sehingga setiap titik ditempati tepat satu orang.  
Alice ingin berdua saja dengan Bob, jadi Alice akan membangun pagar berbentuk persegi panjang dengan posisinya sebagai sudut kiri atas dan posisi Bob sebagai sudut kanan bawah.  
Jika ada orang lain (selain Alice dan Bob) yang berada di dalam atau di atas pagar, maka Alice akan sedih.

Kembalikan jumlah pasangan titik di mana kamu bisa menempatkan Alice dan Bob sehingga Alice tidak sedih saat membangun pagar.

Catatan:  
- Alice hanya bisa membangun pagar dengan posisinya sebagai sudut kiri atas dan Bob sebagai sudut kanan bawah.

**English:**  
You are given a 2D array `points` of size n x 2 representing integer coordinates of some points on a 2D-plane, where `points[i] = [xi, yi]`.

You have to place n people, including Alice and Bob, at these points such that there is exactly one person at every point.  
Alice wants to be alone with Bob, so Alice will build a rectangular fence with Alice's position as the upper left corner and Bob's position as the lower right corner of the fence (the fence might be a line).  
If any person other than Alice and Bob is either inside the fence or on the fence, Alice will be sad.

Return the number of pairs of points where you can place Alice and Bob, such that Alice does not become sad on building the fence.

Note:  
- Alice can only build a fence with her position as the upper left corner and Bob's position as the lower right corner.

---

## Solution

### 1. Greedy Solution

**Indonesia:**  
- Urutkan titik berdasarkan koordinat x (dan y jika x sama).
- Untuk setiap titik A, tentukan batas awal untuk x dan y.
- Untuk setiap titik B setelah A, cek apakah B berada dalam batas yang valid (kanan dan bawah dari A).
- Jika valid, perbarui batas dan tambahkan ke hasil.

**English:**  
- Sort the points by x-coordinate (and by y-coordinate if x is equal).
- For each point A, set initial boundaries for x and y.
- For each point B after A, check if B is within valid boundaries (to the right and below A).
- If valid, update boundaries and increment the answer.

---

### 2. Prefix Sum + Coordinate Compression Solution

**Indonesia:**  
- Karena koordinat bisa sangat besar, lakukan kompresi koordinat agar bisa diproses dalam grid kecil.
- Buat grid dan prefix sum 2D untuk menandai keberadaan titik.
- Untuk setiap pasangan (A, B) yang memenuhi syarat posisi (A.x ≤ B.x dan A.y ≥ B.y), gunakan prefix sum untuk menghitung jumlah titik di dalam persegi panjang yang dibentuk oleh A dan B.
- Jika hanya ada 2 titik (Alice dan Bob) di area tersebut, pasangan valid.

**English:**  
- Since coordinates can be very large, compress the coordinates to process them in a smaller grid.
- Build a grid and 2D prefix sum to mark the presence of points.
- For each pair (A, B) that satisfies the position condition (A.x ≤ B.x and A.y ≥ B.y), use prefix sum to count the number of points inside the rectangle formed by A and B.
- If there are only 2 points (Alice and Bob) in that area, the pair is valid.

---

**Kompleksitas Waktu / Time Complexity:**  
- Greedy: O(n²)  
- Prefix Sum + Compression: O(n² + M²), dengan M adalah jumlah koordinat unik.

**Catatan:**  
- Solusi greedy lebih sederhana untuk n kecil.
- Solusi prefix sum lebih efisien untuk koordinat besar dan n sedang.