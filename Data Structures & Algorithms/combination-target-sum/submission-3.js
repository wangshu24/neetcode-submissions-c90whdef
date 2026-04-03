class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @returns {number[][]}
     */
    combinationSum(nums, target) {
        let res = []
        this.backtrack(nums, target, 0, [], res)

        return res
    }

    backtrack(nums, remain,i, subset, res) {
        if (remain === 0) {
            res.push([...subset])
            return
        } else if (i >= nums.length || remain < 0) {
            return
        }

        subset.push(nums[i])
        this.backtrack(nums, remain - nums[i], i, subset, res)
        subset.pop()
        this.backtrack(nums, remain, i+1, subset, res)
    }
}
