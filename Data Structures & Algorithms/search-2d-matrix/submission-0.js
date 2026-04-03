class Solution {
    /**
     * @param {number[][]} matrix
     * @param {number} target
     * @return {boolean}
     */
    searchMatrix(matrix, target) {
        let flat = matrix.flat()
        let l = 0
        let r = flat.length -1
        
        while(l<=r){
            const mid = Math.trunc((l+r)/2)
            if(flat[mid]=== target){
                return true
            } else if (flat[mid]< target){
                l = mid+1
            } else {
                r = mid-1
            }
        }
        return false
    }
}
