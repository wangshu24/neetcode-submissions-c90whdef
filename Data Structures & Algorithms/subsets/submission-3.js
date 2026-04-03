class Solution {
    /**
     * @param {number[]} nums
     * @return {number[][]}
     */
    subsets(nums) {
        let res = []

        const dfs = (ind, subset) => {
            if (ind === nums.length) {
                res.push([...subset])
                return
            }

            subset.push(nums[ind])
            dfs(ind+1, subset)
            subset.pop()
            dfs(ind+1, subset)
        }
        dfs(0, [])

        return res
    }
}
