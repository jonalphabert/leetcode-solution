## Problem

**Indonesia:**  
Diberikan sebuah linked list, tukar setiap dua node yang bersebelahan dan kembalikan head-nya.  
Kamu harus menyelesaikan masalah ini tanpa mengubah nilai pada node (hanya node itu sendiri yang boleh diubah).

**English:**  
Given a linked list, swap every two adjacent nodes and return its head.  
You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed).

---

## Solution

### Bahasa Indonesia

- Buat node dummy sebagai penunjuk awal.
- Iterasi linked list dua node sekaligus.
- Untuk setiap pasangan node, tukar posisi node dengan mengubah pointer `Next` saja, tanpa mengubah nilai.
- Setelah swap, lanjutkan ke pasangan berikutnya.
- Kembalikan `dummy.Next` sebagai head baru.

### English

- Create a dummy node as the starting pointer.
- Iterate through the linked list two nodes at a time.
- For each pair of nodes, swap their positions by changing the `Next` pointers only, without modifying node values.
- After swapping, move to the next pair.
- Return `dummy.Next` as the new head.

**Time Complexity:** O(n)
**Space Complexity:** O(1)