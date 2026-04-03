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

    backtrack(nums, i, subset, res) {
        if (i === nums.length) {
            res.push([...subset])
        } else if (i > nums.length){
            return
        }

        subset.push(nums[i])
        this.backtrack(nums, i+1, subset, res)
        while (i+1 < nums.length && nums[i] === nums[i+1]) {
            i++
        }
        subset.pop()
        this.backtrack(nums, i+1, subset, res)
    }
}
