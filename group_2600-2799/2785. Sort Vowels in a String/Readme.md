# 📖 2785. Sort Vowels in a String

## 🇮🇩 Bahasa Indonesia

### Deskripsi

Program ini mengimplementasikan fungsi `sortVowels` yang menerima sebuah string `s` lalu menghasilkan string baru `t` dengan aturan:

1. **Konsonan** tetap berada di posisi aslinya.
2. **Vokal** (`a, e, i, o, u` termasuk huruf besar) diurutkan berdasarkan nilai ASCII secara non-decreasing (dari kecil ke besar).

### Contoh

Input:

```
s = "lEetcOde"
```

Output:

```
"lEOtcede"
```

Penjelasan:

* Konsonan `l, t, c, d` tetap di tempat.
* Vokal `E, e, O, e` diurutkan berdasarkan ASCII → `E, O, e, e`.

### Cara Kerja

1. Identifikasi huruf vokal dalam string dan simpan posisinya.
2. Urutkan vokal sesuai nilai ASCII.
3. Masukkan kembali vokal yang sudah diurutkan ke posisi semula.

### Kode Go

```go
func sortVowels(s string) string {
    isVowel := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
    byteArray := []byte(s)
    vowelArray := []byte{}
    posVowel := []int{}


    for i := 0; i < len(s); i++ {
		if isVowel[s[i]] {
			vowelArray = append(vowelArray, s[i])
            posVowel = append(posVowel, i)
		}
	}

    sort.Slice(vowelArray, func(i, j int) bool {
        return int(vowelArray[i]) < int(vowelArray[j]) 
    })

    for index, indexVowel := range posVowel {
        byteArray[indexVowel] = vowelArray[index]
    }

    return string(byteArray)
}
```

---

## 🇬🇧 English

### Description

This program implements a function `sortVowels` that takes a string `s` and produces a new string `t` with the following rules:

1. **Consonants** remain in their original positions.
2. **Vowels** (`a, e, i, o, u` including uppercase) are sorted in non-decreasing ASCII order.

### Example

Input:

```
s = "lEetcOde"
```

Output:

```
"lEOtcede"
```

Explanation:

* Consonants `l, t, c, d` remain unchanged.
* Vowels `E, e, O, e` are sorted by ASCII → `E, O, e, e`.

### How It Works

1. Identify vowels in the string and record their positions.
2. Sort the vowels based on ASCII value.
3. Reinsert the sorted vowels back into their original positions.

### Go Code

```go
func sortVowels(s string) string {
    isVowel := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
    byteArray := []byte(s)
    vowelArray := []byte{}
    posVowel := []int{}


    for i := 0; i < len(s); i++ {
		if isVowel[s[i]] {
			vowelArray = append(vowelArray, s[i])
            posVowel = append(posVowel, i)
		}
	}

    sort.Slice(vowelArray, func(i, j int) bool {
        return int(vowelArray[i]) < int(vowelArray[j]) 
    })

    for index, indexVowel := range posVowel {
        byteArray[indexVowel] = vowelArray[index]
    }

    return string(byteArray)
}
```
