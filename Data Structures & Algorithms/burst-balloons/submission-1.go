

func maxCoins(nums []int) int {
	// Padding: create a new slice with 1s at the boundaries
	// We also filter out 0s as they don't contribute to the score 
	// and don't change the adjacency of other balloons.
	temp := make([]int, 0, len(nums)+2)
	temp = append(temp, 1)
	for _, v := range nums {
		if v > 0 {
			temp = append(temp, v)
		}
	}
	temp = append(temp, 1)
	
	n := len(temp)
	// dp[i][j] represents max coins from bursting all balloons 
	// strictly between index i and j.
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// length is the distance between i and j. 
	// We need at least one balloon between them, so min length is 2.
	for length := 2; length < n; length++ {
		for i := 0; i < n-length; i++ {
			j := i + length
			
			// Try every balloon k between i and j as the last one to burst
			for k := i + 1; k < j; k++ {
				// coins gained by bursting k last in the range (i, j)
				currentBurst := temp[i] * temp[k] * temp[j]
				
				// Total = max coins from left + max coins from right + current burst
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+currentBurst)
			}
		}
	}

	return dp[0][n-1]
}