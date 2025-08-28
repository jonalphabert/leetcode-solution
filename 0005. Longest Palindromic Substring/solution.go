func longestPalindrome(s string) string {
    start := 0
    end := 0

    var expandFromCenter func(left, right int) (int,int);
    expandFromCenter = func(left, right int) (int, int) {
        for left >=0 && right < len(s) && s[left] == s[right] {
            left--
            right++
        }

        return left+1, right-1
    }

    for i := range(len(s)) {
        left1, right1 := expandFromCenter(i, i)
        left2, right2 := expandFromCenter(i, i+1)

        if right1 - left1 > end - start {
            start, end = left1, right1
        }
        if right2 - left2 > end -start {
            start, end = left2, right2
        }
    }

    return s[start:end+1]
}