func evalRPN(tokens []string) int {
    stack := []int{}
    
    for _,i := range tokens {
        switch i {
            case "+" :
                a, b := stack[len(stack)-2], stack[len(stack)-1]
                stack = stack[:len(stack)-2]
                stack = append(stack, a + b)
            case "-" :
                a, b := stack[len(stack)-2], stack[len(stack)-1]
                stack = stack[:len(stack)-2]
                stack = append(stack, a - b)
            case "*" :
                a, b := stack[len(stack)-2], stack[len(stack)-1]
                stack = stack[:len(stack)-2]
                stack = append(stack, a * b)
            case "/" : 
                a, b := stack[len(stack)-2], stack[len(stack)-1]
                stack = stack[:len(stack)-2]
                stack = append(stack, a / b)
            default:
                num, _ := strconv.Atoi(i)
                stack = append(stack, num)
        }
    }

    return stack[0] 
}
