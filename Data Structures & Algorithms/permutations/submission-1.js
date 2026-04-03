class Solution {
    /**
     * @param {number[]} nums
     * @return {number[][]}
     */
    permute(nums) {
        let res = []
        let bool = new Array(nums.length).fill(false)
        this.backtrack(nums, bool, [], res)
        return res
    }

    backtrack(nums, bool, subset, res) {
        if (subset.length === nums.length) {
            res.push([...subset])
            return
        } else {

                for (let i = 0; i < nums.length; i++) {
                    if (!bool[i]) {
                        subset.push(nums[i])
                        bool[i] = true
                        this.backtrack(nums, bool, subset, res)
                        subset.pop()
                        bool[i] = false
                    }
            }
        }
        
    }
}
