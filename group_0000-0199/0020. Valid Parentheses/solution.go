func isValid(s string) bool {
    stack := []byte{}

    for i:=0; i<len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack = append(stack, s[i])
        } else if s[i] == ')' || s[i] == '}' || s[i] == ']' {
            if len(stack) == 0 {
                return false
            }

            top := len(stack) - 1

            if 	(s[i] == ')' && stack[top] =='(') || 
				(s[i] == '}' && stack[top]== '{') || 
				(s[i] == ']' && stack[top]=='[') {
                stack = stack[:top]
            } else {
                return false
            }
        }
    }

    return len(stack) == 0
}