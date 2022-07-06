package sol

type CoinRecord struct {
	start, amount int
}

func changeDFS(amount int, coins []int) int {
	cLen := len(coins)
	cache := make(map[CoinRecord]int)
	var dfs func(start, m int) int
	dfs = func(start, m int) int {
		if m == amount { // found 1
			return 1
		}
		if m > amount {
			return 0
		}
		if start >= cLen {
			return 0
		}
		if val, exists := cache[CoinRecord{start: start, amount: m}]; exists {
			return val
		}
		cache[CoinRecord{start: start, amount: m}] = dfs(start+1, m) + dfs(start, m+coins[start])
		return cache[CoinRecord{start: start, amount: m}]
	}
	return dfs(0, 0)
}
