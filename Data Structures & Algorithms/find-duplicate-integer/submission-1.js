class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    findDuplicate(nums) {
        let map = new Map()
        let res = 0
        for (let i = 0; i < nums.length; i++ )  {
            let val = map.get(nums[i]) ? map.get(nums[i]) : 0
            if ( val > 0) {
                res = nums[i]
                break
            }
            console.log(val)
            map.set(nums[i], val+1)
        }
        return res
    }
}
