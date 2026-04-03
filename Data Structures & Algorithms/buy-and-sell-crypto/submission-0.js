class Solution {
    /**
     * @param {number[]} prices
     * @return {number}
     */
    maxProfit(prices) {
        let res = 0
        let l = 0
        for (let i=1; i < prices.length;i++){
            if (prices[i] < prices[l]){
                l = i
            } else {
                res = Math.max(res, prices[i]-prices[l] )
            }
        }

        return res
    }
}
