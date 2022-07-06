# golang_coin_change_v2

You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money.

Return *the number of combinations that make up that amount*. If that amount of money cannot be made up by any combination of the coins, return `0`.

You may assume that you have an infinite number of each kind of coin.

The answer is **guaranteed** to fit into a signed **32-bit** integer.

## Examples

**Example 1:**

```
Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

```

**Example 2:**

```
Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.

```

**Example 3:**

```
Input: amount = 10, coins = [10]
Output: 1

```

**Constraints:**

- `1 <= coins.length <= 300`
- `1 <= coins[i] <= 5000`
- All the values of `coins` are **unique**.
- `0 <= amount <= 5000`

## 解析

給一個整數 amount 代表要兌換的數目, 還有一個陣列 coins 代表各種幣的面額

要求寫一個演算法來算出能組成該 amount 的方法數

先舉例觀察問題

amount: 5, coins:[1,2,5]

先考慮窮舉法

![](https://i.imgur.com/4mqNDjm.png)

假設 m = len(coins), amount = n

則要走訪完所有結點要 O(m*n)

所時如果直接做 dfs 是 O(m*n)

然而可以發現有些結點其實是重複的

所以可以透過紀錄已經組合過的和來做運算

prune 掉超過的合 總供最多 n 種和

然後每個 和有 m 總可能分支

所以使用 cache 之後 會是 時間複雜度是 O(m*n)

透過動態規劃的作法

先檢視這個問題

定義 dp[amount, start] = 從第 start coin 開始組合成 amount的方法數

dp[amount, start] = dp[amount, start+1]+ dp[amount-coins[start], start]  if amount - coins[start] ≥ 0

dp[amount, start] = dp[amount, start+1]                                                     otherwise

![](https://i.imgur.com/OJHVY5a.png)

一樣 amount 有 0~ n

coins 開始位置有 0~m

所以時間複雜度也是 O(m*n)

空間複雜度是 O(m*n)

另外可以發現從上面關係式知道每次只推算下一個 dp[amount, start] 只關聯到下一個 row 

所以每次只要保持兩個 row 因此空間複雜度其實可以降到 O(n)


## 程式碼
```go
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
```
## 困難點

1. 要能找出coins 起始位置與組成 amount 的遞迴關係式

## Solve Point

- [x]  建立一個 amount+1 by  len(coins）+1 的矩陣 初始化 d[0:] = 1
- [x]  透過遞迴關係式推算由 start = len(coins) -1 遞減往回推算每個 dp 數目
- [x]  回傳 dp[amount][0]