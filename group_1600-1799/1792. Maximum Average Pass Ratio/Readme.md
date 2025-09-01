## Problem

**Indonesia:**  
Terdapat beberapa kelas di sebuah sekolah, masing-masing akan mengikuti ujian akhir. Diberikan array 2D `classes`, di mana `classes[i] = [passi, totali]` berarti pada kelas ke-i terdapat `totali` siswa dan hanya `passi` siswa yang dipastikan lulus ujian.

Diberikan juga integer `extraStudents`, yaitu jumlah siswa tambahan yang dijamin lulus ujian di kelas manapun mereka ditempatkan.  
Tugasnya adalah menempatkan setiap siswa ekstra ke kelas tertentu agar **rata-rata pass ratio** di semua kelas menjadi maksimum.

Pass ratio sebuah kelas adalah jumlah siswa lulus dibagi total siswa di kelas tersebut.  
Rata-rata pass ratio adalah jumlah pass ratio semua kelas dibagi jumlah kelas.

Kembalikan rata-rata pass ratio maksimum setelah semua siswa ekstra ditempatkan.

**English:**  
There is a school with several classes, each having a final exam. You are given a 2D integer array `classes`, where `classes[i] = [passi, totali]` means in the ith class, there are `totali` students and only `passi` students will pass the exam.

You are also given an integer `extraStudents`, representing the number of brilliant students guaranteed to pass the exam in any class they are assigned to.  
Your task is to assign each extra student to a class in a way that **maximizes the average pass ratio** across all classes.

The pass ratio of a class is the number of students passing divided by the total number of students in that class.  
The average pass ratio is the sum of pass ratios of all classes divided by the number of classes.

Return the maximum possible average pass ratio after assigning the extra students.

---

## Algoritma yang biasa dipakai

**Indonesia:**

1. **Hitung gain awal untuk semua kelas.**  
   Gain adalah selisih kenaikan pass ratio jika satu siswa ekstra ditambahkan ke kelas tersebut:
   ```
   gain = (pass + 1) / (total + 1) - pass / total
   ```

2. **Masukkan semua kelas ke dalam max-heap berdasarkan gain.**  
   Heap digunakan agar setiap kali dapat mengambil kelas dengan potensi kenaikan pass ratio terbesar.

3. **Ulangi sebanyak extraStudents kali:**  
   - Ambil kelas dengan gain terbesar dari heap.
   - Tambahkan 1 murid ekstra ke kelas tersebut (pass++, total++).
   - Hitung gain barunya, masukkan lagi ke heap.

4. **Setelah semua extraStudents ditempatkan, hitung average pass ratio:**  
   - Untuk setiap kelas, hitung pass ratio: `pass / total`.
   - Rata-rata = jumlah semua pass ratio dibagi jumlah kelas.

**English:**

1. **Calculate the initial gain for all classes.**  
   Gain is the increase in pass ratio if one extra student is added to the class:
   ```
   gain = (pass + 1) / (total + 1) - pass / total
   ```

2. **Insert all classes into a max-heap based on gain.**  
   The heap allows us to always pick the class with the highest potential increase in pass ratio.

3. **Repeat for extraStudents times:**  
   - Pop the class with the highest gain from the heap.
   - Add 1 extra passing student to that class (pass++, total++).
   - Recalculate its gain and push it back into the heap.

4. **After all extraStudents are assigned, calculate the average pass ratio:**  
   - For each class, calculate pass ratio: `pass / total`.
   - The average is the sum of all pass ratios divided by the number of classes.

**General Pseudocode:**
```bash
function maxAverageRatio(classes, extraStudents):
    # Step 1: definisi gain function
    function gain(pass, total):
        return (pass + 1) / (total + 1) - (pass / total)

    # Step 2: buat max-heap berdasarkan gain
    heap = maxHeap()

    for each (pass, total) in classes:
        push heap with (gain(pass, total), pass, total)

    # Step 3: assign extraStudents satu per satu
    while extraStudents > 0:
        (g, pass, total) = pop from heap   # ambil kelas dengan gain terbesar
        pass = pass + 1
        total = total + 1
        newGain = gain(pass, total)
        push heap with (newGain, pass, total)
        extraStudents = extraStudents - 1

    # Step 4: hitung average pass ratio
    sumRatio = 0
    for each (_, pass, total) in heap:
        sumRatio = sumRatio + (pass / total)

    return sumRatio / number_of_classes
```

