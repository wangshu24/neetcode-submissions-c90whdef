func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _,i := range tokens {
		switch i {
            case "+": 
                b := stack[len(stack)-1]
                a :=  stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack,a+b)
            case "-": 
                b := stack[len(stack)-1]
                a :=  stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack, a-b)
            case "*": 
                b := stack[len(stack)-1]
                a :=  stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack, a*b)
            case "/": 
                b := stack[len(stack)-1] 
                a :=  stack[len(stack)-2]
                stack = stack[:len(stack)-2]
                stack = append(stack, a/b)
            default: 
                num,_ := strconv.Atoi(i)
                stack = append(stack, num)
        }
	}
    fmt.Println(stack)
	return stack[0]
}
