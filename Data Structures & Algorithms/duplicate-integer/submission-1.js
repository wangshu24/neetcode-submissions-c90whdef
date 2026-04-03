class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
        let map = {}

        for (let i=0; i < nums.length; i++) {
            if (!map[nums[i].toString()]) {
                let key = nums[i].toString()
                map[key] =  1

            } else {return true}
        }
    return false
    }
}
