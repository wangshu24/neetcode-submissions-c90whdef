class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @returns {number[][]}
     */
    combinationSum(nums, target) {
        let res = []
        this.dfs(nums, 0, target, [], res)
        return res
    }


    dfs(nums, index,target, subset, res) {
        if (target === 0) {
            res.push([...subset])
        } else if (target < 0 || index >= nums.length) {
            return
        } else {
            subset.push(nums[index])
            this.dfs(nums, index, target - nums[index], subset, res)
            subset.pop()
            this.dfs(nums, index+1, target, subset, res)
        }
    }
}
