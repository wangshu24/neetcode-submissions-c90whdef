class Solution {
    /**
     * @param {string} s
     * @return {number}
     */
    lengthOfLongestSubstring(s) {
        let hset = []
        let res = 0 

        for (const char of s) {
            if (!hset.includes(char)){
                hset.push(char)
            }else {
                while(hset.includes(char)){
                    hset.shift()
                }
                hset.push(char)
            }
            res = Math.max(res, hset.length)
        }

        return res
    }
}
