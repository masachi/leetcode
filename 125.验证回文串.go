/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */
import (
	"fmt";
	"strings";
)
func isPalindrome(s string) bool {

	trimedString := strings.ToLower(strings.Join(regexp.MustCompile(`[0-9]|[a-z]|[A-Z]`).FindAllString(s, -1), ""))
	return trimedString == reverseString(trimedString)
}

func reverseString(s string) string {
    runes := []rune(s)
    for from, to := 0, len(runes)-1; from < to; from, to = from + 1, to - 1 {
        runes[from], runes[to] = runes[to], runes[from] 
    } 

    return string(runes)
}

