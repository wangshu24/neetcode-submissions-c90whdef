class Solution {
    /**
     * @param {number[]} piles
     * @param {number} h
     * @return {number}
     */
    minEatingSpeed(piles, h) {
        let l = 1
        let r = Math.max(...piles)
        let res = r

        while (l <= r){
            const mid = Math.trunc((l+r)/2)
            let times = 0
            for (const pile of piles){
                times += Math.ceil(pile/mid)
            }
            if (times <= h){
                res = mid
                r = mid-1
            } else  {
                
                l = mid+1
            }
            
        }
        return res
    }
}
