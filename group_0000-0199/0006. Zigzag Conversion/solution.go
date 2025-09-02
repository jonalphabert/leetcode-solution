/*
	Solution with pattern observation
*/
func convert(s string, n int) string {
    result := ""
    length := len(s)


    if n == 1 {
        return s
    }

    for i:=0 ; i < n && i < length ; i++ {
        result += string(s[i])

        j := i
        for j < length {
            jump := (n-1-i) * 2

            if j + jump >= length {
                break
            }
            
            if jump != 0 {
                j += jump
                result += string(s[j])
            }

            jump2 := i * 2

            if j + jump2 >= length {
                break
            }

            
            if jump2 != 0{
                j += jump2
                result += string(s[j])
            }
        }
    }

    return result
}

/*
	Solution with simulation to build the zigzag pattern
*/
func convert(s string, n int) string {
    if n == 1 || n >= len(s) {
		return s
	}
    
    result := ""
    rowsBuilder := make([]string, n)
    isGoingDown := false
    currentRow := 0


    for i := 0; i<len(s); i++ {
        rowsBuilder[currentRow] += string(s[i])

        if currentRow == 0 || currentRow == n-1 {
            isGoingDown = !isGoingDown
        }

        if isGoingDown {
            currentRow ++
        } else {
            currentRow --
        }
    }

    for _, val := range rowsBuilder {
        result += string(val)
    }

    return result
}