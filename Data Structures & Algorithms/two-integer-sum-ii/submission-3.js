class Solution {
    /**
     * @param {number[]} numbers
     * @param {number} target
     * @return {number[]}
     */
    twoSum(numbers, target) {
        // let tarIndex = numbers.indexOf(target) === -1 ? numbers.length-1 : numbers.indexOf(target)
        let i = 0
        let j = numbers.length-1
        console.log(i, j)
        let res = []
        while (i < j){
            console.log(numbers[i])
            console.log(numbers[j])
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
        console.log(res)
        return res
    }
}
