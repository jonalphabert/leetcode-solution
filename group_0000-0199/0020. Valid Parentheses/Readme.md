# 📖 0020. Valid Parentheses

## 🇮🇩 Bahasa Indonesia

### Deskripsi

Program ini mengimplementasikan fungsi `isValid` yang memeriksa apakah sebuah string `s` yang hanya berisi karakter `'(' , ')' , '{' , '}' , '[' , ']'` merupakan **urutan tanda kurung yang valid**.

Sebuah string dianggap valid jika:

1. Setiap kurung buka memiliki kurung tutup dengan jenis yang sama.
2. Kurung buka harus ditutup dalam urutan yang benar.
3. Tidak ada kurung tutup yang muncul tanpa kurung buka pasangannya.

### Contoh

Input:

```
s = "()[]{}"
```

Output:

```
true
```

Input:

```
s = "(]"
```

Output:

```
false
```

Input:

```
s = "([)]"
```

Output:

```
false
```

Input:

```
s = "{[]}"
```

Output:

```
true
```

### Cara Kerja

1. Gunakan **stack** untuk menyimpan setiap kurung buka.
2. Setiap kali menemukan kurung tutup, periksa elemen paling atas stack:

   * Jika cocok dengan pasangannya → keluarkan dari stack.
   * Jika tidak cocok → string tidak valid.
3. Jika setelah semua iterasi stack kosong → string valid.

### Kode Go

```go
func isValid(s string) bool {
    stack := []byte{}

    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack = append(stack, s[i])
        } else if s[i] == ')' || s[i] == '}' || s[i] == ']' {
            if len(stack) == 0 {
                return false
            }

            top := len(stack) - 1
            if (s[i] == ')' && stack[top] == '(') ||
               (s[i] == '}' && stack[top] == '{') ||
               (s[i] == ']' && stack[top] == '[') {
                stack = stack[:top] // pop
            } else {
                return false
            }
        }
    }

    return len(stack) == 0
}
```

---

## 🇬🇧 English

### Description

This program implements the `isValid` function to check whether a string `s`, consisting only of `'(' , ')' , '{' , '}' , '[' , ']'`, is a **valid sequence of parentheses**.

A string is considered valid if:

1. Every opening bracket has a corresponding closing bracket of the same type.
2. Opening brackets must be closed in the correct order.
3. No closing bracket appears without its corresponding opening bracket.

### Example

Input:

```
s = "()[]{}"
```

Output:

```
true
```

Input:

```
s = "(]"
```

Output:

```
false
```

Input:

```
s = "([)]"
```

Output:

```
false
```

Input:

```
s = "{[]}"
```

Output:

```
true
```

### How It Works

1. Use a **stack** to store opening brackets.
2. For each closing bracket:

   * If it matches the top of the stack → pop it.
   * Otherwise → invalid string.
3. After processing all characters, if the stack is empty → valid string.

### Go Code

```go
func isValid(s string) bool {
    stack := []byte{}

    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack = append(stack, s[i])
        } else if s[i] == ')' || s[i] == '}' || s[i] == ']' {
            if len(stack) == 0 {
                return false
            }

            top := len(stack) - 1
            if (s[i] == ')' && stack[top] == '(') ||
               (s[i] == '}' && stack[top] == '{') ||
               (s[i] == ']' && stack[top] == '[') {
                stack = stack[:top] // pop
            } else {
                return false
            }
        }
    }

    return len(stack) == 0
}
```