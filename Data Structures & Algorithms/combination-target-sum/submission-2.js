class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @returns {number[][]}
     */
    combinationSum(nums, target) {
        let res = []
        this.backtrack(nums, 0, target, [], res)
        return res
    }

    backtrack(nums, i, remain, subset, res) {
        if (remain === 0) {
            res.push([...subset])
            return
        } else if (i >= nums.length || remain < 0) {
            return
        }

        subset.push(nums[i])
        this.backtrack(nums, i, remain - nums[i], subset, res)
        subset.pop()
        this.backtrack(nums, i+1, remain, subset, res)
    }
}
