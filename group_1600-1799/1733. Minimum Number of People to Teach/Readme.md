## 🇮🇩 Bahasa Indonesia

### 📌 Deskripsi Masalah

Kita memiliki `n` bahasa, `m` pengguna, dan daftar `friendships` (pertemanan).

* Setiap pengguna bisa mengetahui lebih dari satu bahasa.
* Dua pengguna bisa berkomunikasi jika mereka memiliki setidaknya satu bahasa yang sama.
* Kita dapat memilih **satu bahasa** untuk diajarkan ke beberapa pengguna.

Tujuan: cari **jumlah minimum pengguna** yang harus diajarkan bahasa baru agar semua pasangan teman (`friendship`) dapat berkomunikasi.

---

### 🔍 Abstraksi & Ide

1. **Identifikasi edge gagal komunikasi**

   * Periksa setiap pasangan teman `(u, v)`.
   * Jika mereka tidak memiliki bahasa yang sama → catat sebagai gagal.

2. **Kumpulkan kandidat pengguna**

   * Semua user yang muncul di edge gagal masuk ke kandidat.

3. **Hitung frekuensi bahasa**

   * Untuk setiap bahasa `L`, hitung berapa kandidat yang sudah bisa `L`.
   * Jika kita pilih `L`, maka hanya perlu mengajarkan ke `total_kandidat - count(L)`.

4. **Ambil minimum global**

   * Jawaban adalah minimum dari semua bahasa.

---

### 💡 Penyelesaian

* Gunakan **hash set** untuk mengecek apakah dua pengguna memiliki bahasa yang sama dengan efisien (`O(min(len(u), len(v)))`).
* Gunakan **map/set** untuk menyimpan semua user gagal komunikasi.
* Hitung **frekuensi bahasa** di antara user gagal → pilih bahasa yang paling sering dipakai.
* Hasil = `len(user_gagal) - max_freq`.

---

### ⏱️ Kompleksitas

* **Waktu:**

  * Pemeriksaan semua `friendships`: `O(len(friendships) * avg_lang_per_user)`
  * Hitung frekuensi bahasa: `O(sum(len(languages[u]) for u in failedUsers))`
  * Total: `O(len(friendships) + n)` dalam praktik.

* **Ruang:**

  * Penyimpanan bahasa tiap user + counter bahasa = `O(n + m)`.

---

## 🇬🇧 English

### 📌 Problem Description

We are given `n` languages, `m` users, and a list of `friendships`.

* Each user may know multiple languages.
* Two users can communicate if they share at least one common language.
* We can choose **one language** to teach to some users.

Goal: find the **minimum number of users** that need to learn the chosen language so that all friendship pairs can communicate.

---

### 🔍 Abstraction & Idea

1. **Identify failed communication edges**

   * For each friendship `(u, v)`, check if they share a language.
   * If not, mark both users as “failed”.

2. **Collect candidate users**

   * All users involved in failed friendships are candidates.

3. **Count language frequency**

   * For each language `L`, count how many candidates already know it.
   * If we choose `L`, then we need to teach `total_candidates - count(L)`.

4. **Take the global minimum**

   * Answer = minimum across all languages.

---

### 💡 Solution

* Use **hash sets** to check whether two users share a language efficiently (`O(min(len(u), len(v)))`).
* Maintain a **set of failed users** that appear in failed friendships.
* Compute **language frequencies** among failed users → choose the most common one.
* Result = `len(failedUsers) - maxFrequency`.

---

### ⏱️ Complexity

* **Time:**

  * Check all friendships: `O(len(friendships) * avg_lang_per_user)`
  * Count frequencies: `O(sum(len(languages[u]) for u in failedUsers))`
  * Overall ≈ `O(len(friendships) + n)` in practice.

* **Space:**

  * Store user languages + counters = `O(n + m)`.

---
