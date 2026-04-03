class Solution {
    /**
     * @param {string} s1
     * @param {string} s2
     * @return {boolean}
     */
    checkInclusion(s1, s2) {
        let s1arr = new Array(26).fill(0)
        let s2arr = new Array(26).fill(0)


        for (let i=0; i < s1.length;i++) {
            s1arr[s1.charCodeAt(i)-97]++
            s2arr[s2.charCodeAt(i)-97]++
        }

        let matches = 0
        for (let i=0; i < s1arr.length;i++) {
            if (s1arr[i] === s2arr[i]){
                matches++
            }
        }
        
        let l = 0

        for(let i = s1.length; i <= s2.length; i++){
            if(matches === 26) {return true}

            let index = s2.charCodeAt(i)-97 
            s2arr[index]++
            if (s2arr[index] === s1arr[index]) {
                matches++
            } else if(s2arr[index] - 1 === s1arr[index])  {
                matches--;
            }

            index = s2.charCodeAt(l)-97
            s2arr[index]--
            if(s2arr[index] === s1arr[index]){
                matches++
            } else if(s2arr[index] +1 === s1arr[index]) {
                matches--;
            }
            
            l++
            if (matches === 26){return true}
        }

        return matches===26
    }
}
