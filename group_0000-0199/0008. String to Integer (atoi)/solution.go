import (
    "strings"
)

func myAtoi(s string) int {
    const INT_MAX = 1<<31 - 1
    const INT_MIN = -1 << 31

    s = strings.TrimSpace(s)

    multiplier := 1
    index := 0
    res := 0

    if index < len(s) && (s[index] == '-' || s[index] == '+') {
        if s[index] == '-' {
            multiplier = -1
        }
        index++
    }

    for index < len(s) && s[index] >= '0' && s[index] <= '9' {
        digit := int(s[index] - '0')

        // check overflow sebelum masukin digit baru
        if res > (INT_MAX-digit)/10 {
            if multiplier == 1 {
                return INT_MAX
            }
            return INT_MIN
        }

        res = res*10 + digit
        index++
    }

    return multiplier * res
}