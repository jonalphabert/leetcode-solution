# Problem

## English
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

## Indonesia
Diberikan dua buah array yang telah diurutkan sebelumnya dengan num1 dan nums2 memiliki panjang m dan n. Kembalikan nilai median dari kedua array tersebut.

Untuk kompleksitas waktu solusi harus pada O(log (m+n)).

# Testcase
Example 1:

```
    Input: nums1 = [1,3], nums2 = [2]
    Output: 2.00000
    Explanation: merged array = [1,2,3] and median is 2.
```

Example 2:
```
    Input: nums1 = [1,2], nums2 = [3,4]
    Output: 2.50000
    Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
```

# 📌 Intuisi Penyelesaian

### Median definisi
    Jika total elemen (m+n) ganjil → median adalah elemen ke-(m+n)//2.
    Jika genap → median adalah rata-rata dari elemen ke-(m+n)//2 - 1 dan (m+n)//2.

### Binary Search Partition
    Kita bisa membayangkan ada “pembatas” (partition) pada kedua array.
    Misalnya, array nums1 dibagi menjadi dua bagian: kiri dan kanan.
    Array nums2 juga dibagi dua.
    Tujuan: bagian kiri gabungan berisi setengah elemen (atau setengah+1 kalau total ganjil), dan semua elemen di kiri ≤ semua elemen di kanan.

### Binary search

    Kita selalu lakukan binary search di array yang lebih kecil (nums1 misalnya).
    Coba letakkan i elemen di nums1 untuk bagian kiri, maka j = half - i elemen diambil dari nums2.

    Cek apakah pembatas valid:

        nums1[i-1] <= nums2[j] 
    dan
        nums2[j-1] <= nums1[i]


    Jika belum valid, geser binary search.

### Hasil median
    Jika total (m+n) ganjil → median adalah max(left_part).
    Jika genap → median adalah (max(left_part) + min(right_part)) / 2.