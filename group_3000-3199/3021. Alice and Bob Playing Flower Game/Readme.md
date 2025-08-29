## Problem

**Indonesia:**  
Alice dan Bob bermain game secara bergiliran di sebuah lapangan dengan dua jalur bunga di antara mereka. Terdapat `x` bunga di jalur pertama dan `y` bunga di jalur kedua. Permainan berlangsung sebagai berikut: Alice selalu jalan dulu. Pada setiap giliran, pemain harus memilih salah satu jalur dan mengambil satu bunga dari jalur tersebut. Jika setelah giliran tidak ada bunga tersisa sama sekali, pemain yang mengambil bunga terakhir langsung menang dengan menangkap lawannya.  
Diberikan dua bilangan bulat `n` dan `m`, hitung berapa banyak pasangan `(x, y)` yang memenuhi syarat agar Alice pasti menang. Nilai `x` berada pada rentang `[1, n]` dan nilai `y` pada rentang `[1, m]`.  
Kembalikan jumlah pasangan `(x, y)` yang memenuhi syarat tersebut.

**English:**  
Alice and Bob are playing a turn-based game on a field, with two lanes of flowers between them. There are `x` flowers in the first lane and `y` flowers in the second lane. The game proceeds as follows: Alice takes the first turn. In each turn, a player must choose either lane and pick one flower from that side. If, after a turn, there are no flowers left at all, the current player captures their opponent and wins the game.  
Given two integers `n` and `m`, compute the number of possible pairs `(x, y)` such that Alice is guaranteed to win the game. The number of flowers `x` in the first lane must be in the range `[1, n]`, and `y` in the second lane must be in the range `[1, m]`.  
Return the number of possible pairs `(x, y)` that satisfy these conditions.

---

## Solution

### Bahasa Indonesia

#### 1. Aturan Permainan

- Alice jalan dulu.
- Tiap giliran, pemain **ambil 1 bunga dari salah satu jalur**.
- Pemain yang **mengambil bunga terakhir** langsung menang.

Permainan ini identik dengan "take-away game" pada dua tumpukan, di mana setiap langkah hanya boleh mengurangi salah satu tumpukan sebanyak 1.

#### 2. Siapa yang menang?

Jika total bunga $T = x + y$:
- Alice mulai dulu.
- Setiap giliran mengurangi total bunga tepat 1.
- Jumlah langkah = $T$.
- Alice menang jika dia mendapat giliran terakhir, yaitu jika $T$ ganjil.
- Bob menang jika $T$ genap.

#### 3. Kesimpulan

Alice menang **jika dan hanya jika $x + y$ ganjil**, dengan $1 \leq x \leq n$, $1 \leq y \leq m$.

#### 4. Perhitungan pasangan $(x, y)$

- Banyak bilangan ganjil dalam $[1, n]$ = $\lceil n/2 \rceil$.
- Banyak bilangan genap dalam $[1, n]$ = $\lfloor n/2 \rfloor$.
- Banyak bilangan ganjil dalam $[1, m]$ = $\lceil m/2 \rceil$.
- Banyak bilangan genap dalam $[1, m]$ = $\lfloor m/2 \rfloor$.

Supaya $x + y$ ganjil:
- $x$ ganjil & $y$ genap, atau
- $x$ genap & $y$ ganjil.

Jumlah pasangan yang valid:
$$
\text{ans} = \big(\lceil n/2 \rceil \cdot \lfloor m/2 \rfloor\big) + \big(\lfloor n/2 \rfloor \cdot \lceil m/2 \rceil\big)
$$

---

### English

#### 1. Game Rules

- Alice goes first.
- On each turn, a player picks **one flower from either lane**.
- The player who picks the last flower wins immediately.

This game is equivalent to a "take-away game" with two piles, where each move removes exactly one from either pile.

#### 2. Who wins?

If the total number of flowers is $T = x + y$:
- Alice starts first.
- Each turn reduces the total flower count by exactly 1.
- Number of moves = $T$.
- Alice wins if she gets the last move, i.e., if $T$ is odd.
- Bob wins if $T$ is even.

#### 3. Conclusion

Alice wins **if and only if $x + y$ is odd**, with $1 \leq x \leq n$, $1 \leq y \leq m$.

#### 4. Counting valid pairs $(x, y)$

- Number of odd numbers in $[1, n]$ = $\lceil n/2 \rceil$.
- Number of even numbers in $[1, n]$ = $\lfloor n/2 \rfloor$.
- Number of odd numbers in $[1, m]$ = $\lceil m/2 \rceil$.
- Number of even numbers in $[1, m]$ = $\lfloor m/2 \rfloor$.

To make $x + y$ odd:
- $x$ odd & $y$ even, or
- $x$ even & $y$ odd.

Total valid pairs:
$$
\text{ans} = \big(\lceil n/2 \rceil \cdot \lfloor m/2 \rfloor\big) + \big(\lfloor n/2 \rfloor \cdot \lceil m/2 \rceil\big)
$$
