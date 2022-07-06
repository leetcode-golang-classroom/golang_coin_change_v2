package sol

func changeV1(amount int, coins []int) int {
	cLen := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1
	for start := cLen - 1; start >= 0; start-- {
		nextDp := make([]int, amount+1)
		nextDp[0] = 1
		for m := 1; m <= amount; m++ {
			nextDp[m] = dp[m]
			if m-coins[start] >= 0 {
				nextDp[m] += nextDp[m-coins[start]]
			}
		}
		copy(dp, nextDp)
	}
	return dp[amount]
}
