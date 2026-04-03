class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
        let res =  []
        let len = nums.length
        for (let i =0; i < nums.length;i++){
            let remain = target - nums[i]
            let tempNums = nums.slice(i+1,len)
           
            if (tempNums.findIndex((ele) => ele === remain)>-1){
                let j = tempNums.findIndex((ele) => ele === remain) + i + 1
         
                res.push(i)
                res.push(j)
                return res
            }
        }
        return res
    }
}
