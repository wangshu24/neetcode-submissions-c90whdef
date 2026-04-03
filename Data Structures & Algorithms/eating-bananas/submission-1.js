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
        while (l <= r) {
            const mid = Math.trunc((l+r)/2)
            
            let times = 0
            for (let i=0; i < piles.length; i++) {
                times+=Math.ceil(piles[i]/mid)
            }
            console.log(times)
            if(times <= h){
               res = mid
               r = mid -1
            } else {
                l = mid + 1
            }
        }  

        return res
    }
}
