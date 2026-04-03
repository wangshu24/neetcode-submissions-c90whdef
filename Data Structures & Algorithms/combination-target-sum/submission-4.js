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

    backtrack(nums, target,i,subset, res) {
        if (target === 0) {
            res.push([...subset])
            return
        } else if (i >= nums.length || target < 0) {
            return
        }

        subset.push(nums[i])
        this.backtrack(nums, target-nums[i], i, subset,res)
        subset.pop()
        this.backtrack(nums, target, i+1, subset,res)
    }
}
