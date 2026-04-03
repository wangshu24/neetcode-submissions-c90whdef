class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
   

    maxArea(heights) {
        let l = 0
        let r = heights.length -1
        let res = 0
        while (l < r) {
            res = heights[l] > heights[r] ? Math.max(res, heights[r] * (r-l)) : Math.max(res, heights[l] * (r-l))
            if (heights[l] > heights[r]) {
                r--
            } else {l++}
        }
        return res
    }
}
