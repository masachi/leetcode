/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
import (
	"strconv";
	"fmt"
)
func reverse(x int) int {
	inputString := strconv.Itoa(x);
	resultString := ""
	if string(inputString[0]) == string("-") {
		resultString = resultString + string("-")
		resultString = resultString + reverseString(inputString[1:])
	} else {
		resultString = reverseString(inputString)
	}

	result, _ := strconv.Atoi(resultString)

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}

func reverseString(input string) string {
	runes := []rune(input)
	for from, to := 0, len(runes)-1; from < to; from, to = from + 1, to -1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

