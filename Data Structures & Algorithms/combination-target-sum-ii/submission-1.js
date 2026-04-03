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

    backtrack(cand, target, ind, subset, res) {
        if (target === 0) {
            res.push([...subset])
            return
        } else if (ind >= cand.length || target < 0) {
            return
        } else {
            subset.push(cand[ind])
            this.backtrack(cand, target - cand[ind], ind+1,subset, res)
            
            subset.pop()
            while (ind+1 < cand.length && cand[ind] === cand[ind+1]) {
                ind++
            }
            this.backtrack(cand, target, ind+1, subset, res)
        }
    }
}
