/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
 import (
	"strconv"
)
func isPalindrome(x int) bool {
    return strconv.Itoa(x) == reverseString(strconv.Itoa(x))
}

func reverseString(input string) string {
	runes := []rune(input)
	for from, to := 0, len(runes)-1; from < to; from, to = from + 1, to -1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

