class Solution {
    /**
     * @param {number[]} candidates
     * @param {number} target
     * @return {number[][]}
     */
    combinationSum2(candidates, target) {
        let res = []
        candidates.sort()
        this.backtrack(candidates, target, 0, [], res)
        return res    
}

    backtrack(cand, target, i, subset, res) {
        if (target === 0) {
            res.push([...subset])
            return res
        } else if (target < 0 || i >= cand.length) {
            return
        }

        subset.push(cand[i])
        this.backtrack(cand, target-cand[i], i+1, subset, res)
        
        subset.pop()
        while (i+1 <cand.length && cand[i] == cand[i+1]) {i++}
        this.backtrack(cand, target, i+1, subset, res)
    }
}
