class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isValid(s) {
        let que = []
        
        if(s.length === 1) {return false}

        for (let i=0; i < s.length; i++){
            if (["(", "{","["].includes(s[i])){
                que.push(s[i])
            }  else {
                switch(s[i]){
                    case ")":
                        if(que[que.length-1] !== "("){
                            return false
                        }
                        que.pop()
                        break
                    case "}":
                        if(que[que.length-1] !== "{"){
                            return false
                        }
                        que.pop()
                        break
                    default:
                        if(que[que.length-1] !== "["){
                            return false
                        }
                        que.pop()
                }
            }      
            }
        if (que.length > 0) {return false}
        return true
    }
}
