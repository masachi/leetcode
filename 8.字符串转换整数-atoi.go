/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */
import (
	"strings";
	"regexp";
	"strconv"
)

func myAtoi(str string) int {
	trimedStr := strings.Trim(str, " ")

	result := 0

    if len(trimedStr) == 0 {
		return result
	}

	if len(trimedStr) == 1 {
		if string(trimedStr[0]) == string("-") || string(trimedStr[0]) == string("+") {
			return result
		}
	}

	if string(trimedStr[0]) == string("-") || string(trimedStr[0]) == string("+") {
		matched, _ := regexp.MatchString("[0-9]", string(trimedStr[1]))
		if matched {
			result, _ = strconv.Atoi(regexp.MustCompile(`\d+`).FindAllString(trimedStr, 1)[0])
			if string(trimedStr[0]) == string("-") {
				resultStr := string("-") + strconv.Itoa(result)
				result, _ = strconv.Atoi(resultStr)
			}
		}
	}

	matched, _ := regexp.MatchString("[0-9]", string(trimedStr[0]))
	
	if matched {
		result, _ = strconv.Atoi(regexp.MustCompile(`\d+`).FindAllString(trimedStr, 1)[0])
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}

	if result < math.MinInt32  {
		return math.MinInt32
	}

	return result
}

