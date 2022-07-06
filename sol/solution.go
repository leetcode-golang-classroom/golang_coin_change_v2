package sol

func change(amount int, coins []int) int {
	cLen := len(coins)
	dp := make([][]int, amount+1)
	for row := range dp {
		dp[row] = make([]int, cLen+1)
	}
	// init dp[0][coins] = 1
	for start := 0; start <= cLen; start++ {
		dp[0][start] = 1
	}
	for m := 1; m <= amount; m++ {
		for start := cLen - 1; start >= 0; start-- {
			dp[m][start] = dp[m][start+1]
			if m-coins[start] >= 0 {
				dp[m][start] += dp[m-coins[start]][start]
			}
		}
	}
	return dp[amount][0]
}
