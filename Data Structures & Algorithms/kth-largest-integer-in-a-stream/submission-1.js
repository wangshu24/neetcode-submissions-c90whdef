class KthLargest {
    /**
     * @param {number} k
     * @param {number[]} nums
     */
    constructor(k, nums) {
        this.k = k
        this.nums = [...nums].sort((a,b) => b-a)
    }

    /**
     * @param {number} val
     * @return {number}
     */
    add(val) {
        this.nums.unshift(val)
        this.nums.sort((a,b) => b-a)
        return this.nums[this.k-1]
    }
}
