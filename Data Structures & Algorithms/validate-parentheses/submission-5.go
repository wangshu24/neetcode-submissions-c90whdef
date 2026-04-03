func isValid(s string) bool {
    // for _,i := range s {
    //     fmt.Println(i)
    // }
    
    stack := []byte{}
    for i:=0; i < len(s); i++ {
        char := s[i]
        fmt.Println(char)
        if char == 41 {
            if len(stack) > 0 && stack[len(stack)-1] == 40 {
                stack = stack[:len(stack)-1]
            } else {return false}
        } else if char == 93 {
            if len(stack) > 0 && stack[len(stack)-1] == 91 {
                stack=stack[:len(stack)-1]
            } else {return false}
        } else if  len(stack) > 0 && char == 125 {
            if stack[len(stack)-1] == 123 {
                stack = stack[:len(stack)-1]
            } else {return false}
        } else {
            stack = append(stack, char)
        }
    }
    res := len(stack) > 0
    return !res
}
