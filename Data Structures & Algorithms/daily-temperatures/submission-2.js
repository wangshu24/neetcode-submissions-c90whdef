class Solution {
    /**
     * @param {number[]} temperatures
     * @return {number[]}
     */
    dailyTemperatures(temp) {
        let res = []

        for (let i=0; i < temp.length;i++){
            // if (i+1 === temp.length){
            //     res.push(0)
            //     continue
            // }
            const cur = temp[i]
            let j = i+1
            let count = 1
            while(temp[j] <= cur && j < temp.length){
                j++
                count++
            }

            count = j === temp.length ?  0 : count

            res.push(count)

        }

        return res
    }
}
