class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {
        let map = {}
         for (let str of strs) {
            const key = str.split('').sort().join('')
            if (!map[key]){
                const val = [str]
                map[key] = val
            } else {
                map[key].push(str)
            }
         }
        let res = []
         for (let obj in map){
         
            res.push(map[obj])
         }
         return res
    }
}
