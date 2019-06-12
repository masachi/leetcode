/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	if len(s) == 0 {
		return true
	}

	if len(s) % 2 == 1 {
		return false
	}

    for index := 0; index < len(s) -1; index++ {
		if string(s[index]) == string("(") {
			if string(s[index + 1]) != string(")") {
				return false
			}
		}

		if string(s[index]) == string("[") {
			if string(s[index + 1]) != string("]") {
				return false
			}
		}

		if string(s[index]) == string("{") {
			if string(s[index + 1]) != string("}") {
				return false
			}
		}
	}

	return true
}

