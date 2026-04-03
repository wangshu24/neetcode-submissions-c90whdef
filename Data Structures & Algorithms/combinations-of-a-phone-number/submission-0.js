class Solution {
    /**
     * @param {string} digits
     * @return {string[]}
     */
    letterCombinations(digits) {
        if (digits.length === 0) {return []}
        const hash = {
            "2" : ["a","b","c"],
            "3" : ["d", "e", "f"],
            "4" : ["g", "h", "i"],
            "5" : ["j", "k", "l"],
            "6" : ["m", "n", "o"],
            "7" : ["p", "q", "r", "s"],
            "8" : ["t", "u", "v"],
            "9" : ["w", "x", "y", "z"]
        }
        let res = []
        this.backtrack(0, digits, hash, [], res)
        return res
    }

    backtrack(i, digits, hash, subset, res) {
        if (i === digits.length) {
            res.push([...subset].join(""))
            return
        }

        const opts = hash[digits[i]]
        console.log(opts)
        for (const opt of opts) {
            subset.push(opt)
            this.backtrack(i+1, digits, hash, subset, res)
            subset.pop()
        }
    }
}
