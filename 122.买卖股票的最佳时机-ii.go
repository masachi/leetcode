/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */
func maxProfit(prices []int) int {
    count := len(prices)

	if count == 0 {
		return 0
	}

	// var dp [30000][2]int

	result := 0

	for index := 1; index < count; index++ {
		// if index - 1 == -1 {
		// 	dp[index][0] = 0;
		// 	dp[index][1] = math.MinInt32;
		// 	continue;
		// }
		// temp := dp[index][0]
		// dp[index][0] = int(math.Max(float64(dp[index - 1][0]), float64(dp[index - 1][1] + prices[index])))
		// dp[index][1] = int(math.Max(float64(dp[index - 1][1]), float64(temp - prices[index])))

		if prices[index] > prices[index - 1] {
			result = result + prices[index] - prices[index -1] 
		}
	}

	return result
}

