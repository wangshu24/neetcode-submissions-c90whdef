class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {
        let map1={}
        let map2={}
        if (s.length != t.length){return false}
        for (let i=0; i < s.length; i++) {
            if (!map1[s[i]]) {
                map1[s[i]] = 1
            } else {map1[s[i]]++}
            if (!map2[t[i]]) {
                map2[t[i]] = 1
            } else {map2[t[i]]++}
        }

        let keys1 = Object.keys(map1)
        for (let i=0; i < keys1.length;i++) {
            if (map1[keys1[i]] != map2[keys1[i]]) {return false}
        }

        return true
    }
}
