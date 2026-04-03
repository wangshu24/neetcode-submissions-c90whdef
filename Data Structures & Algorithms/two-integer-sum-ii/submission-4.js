class Solution {
    /**
     * @param {number[]} numbers
     * @param {number} target
     * @return {number[]}
     */
    twoSum(numbers, target) {
        let i = 0
        let j = numbers.length-1
        let res = []
        while (i < j){
            if ((numbers[i] + numbers[j]) === target) {
                res.push(i+1)
                res.push(j+1)
                break
            } else {
                if ((numbers[i] + numbers[j]) > target){
                    j--
                }else {i++}
            }
        }
        return res
    }
}
