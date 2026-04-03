class Solution {
    /**
     * @param {string[]} tokens
     * @return {number}
     */
    evalRPN(tokens) {
        let que = []
        
    for (const token of tokens){
        if (token === "+"){
            que.push(que.pop()+que.pop())
        } else if (token === "-") {
            const a = que.pop()
            const b = que.pop()
            que.push(b-a)
        } else if (token === "*"){
            const a = que.pop()
            const b = que.pop()
            que.push(a*b)
        } else if (token === "/"){
            const a = que.pop()
            const b = que.pop()
            console.log(a, b, b/a)
            que.push(Math.trunc(b/a))
        } else {
            que.push(parseInt(token))
        }
    }

        return que.pop()
    }
}
