class Solution {
    /**
     * @param {number[]} nums
     * @return {number[][]}
     */
    subsetsWithDup(nums) {
        let res = []
        nums.sort()
        this.backtrack(nums, 0, [], res)
        return res
    }

    backtrack(nums, ind, subset, res) {
        if (ind >= nums.length) {
            res.push([...subset])
            return
        }

        subset.push(nums[ind])
        this.backtrack(nums, ind+1, subset, res)
        subset.pop()

        while (ind+1 < nums.length && nums[ind] === nums[ind+1]){
            ind++
        }
        this.backtrack(nums, ind+1, subset, res)

    }
}
