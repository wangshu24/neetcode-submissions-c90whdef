class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
        let map = {}
        for (let num of nums){
            if(!map[num]){
                map[num] = 1
            } else {map[num]++}
        }
        let res = []
        
        let arrMap = Object.entries(map)
        arrMap.sort((arrA, arrB) => {return arrA[1] - arrB[1]})
        console.log(arrMap)
        for (let i=0; i < k;i++) {
            res.push(arrMap[arrMap.length-1-i][0])
        }
        return res
    }
}
