package main

import "strconv"

//设计一种方法，通过给重复字符计数来进行基本的字符串压缩。
//例如，字符串 aabcccccaaa 可压缩为 a2b1c5a3 。而如果压缩后的字符数不小于原始的字符数，则返回原始的字符串。
//可以假设字符串仅包括 a-z 的大/小写字母。

/**
 * @param originalString: a string
 * @return: a compressed string
 */
func compress(originalString string) string {
	if originalString == "" {
		return ""
	}

	var (
		prev   string
		count  int
		result string
	)

	for _, o := range originalString {
		now := string(o)
		if prev != "" && prev != now {
			result += prev + strconv.Itoa(count)
			prev = now
			count = 1
		} else {
			prev = now
			count++
		}
	}
	result += prev + strconv.Itoa(count)

	if len(result) >= len(originalString) {
		return originalString
	}
	return result
}
