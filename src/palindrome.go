package src

/* 给你一个字符串 s，找到 s 中最长的回文子串 */
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		// 回文串长度为奇数
		p1 := palindrome(s, i, i)
		if len(p1) > len(res) {
			res = p1
		}

		// 回文串长度为偶数
		p2 := palindrome(s, i, i+1)
		if len(p2) > len(res) {
			res = p2
		}
	}

	return res
}

// 从中心向两段扩展
func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	// (r-1)-l 为回文串的长度
	// s[l+1 : (r-1)-l+(l+1)] 为回文串
	// 区间：[:)
	return s[l+1 : r]
}
