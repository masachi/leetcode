/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
func maxProfit(prices []int) int {
	count := len(prices)

	if count == 0 {
		return 0
	}

	var dp [30000][2]int
	for index := 0; index < count; index++ {
		if index - 1 == -1 {
			dp[index][0] = 0;
			dp[index][1] = -prices[index];
			continue;
		}

		dp[index][0] = int(math.Max(float64(dp[index - 1][0]), float64(dp[index - 1][1] + prices[index])))
		dp[index][1] = int(math.Max(float64(dp[index - 1][1]), float64(-prices[index])))
	}

	return dp[count - 1][0]
}

