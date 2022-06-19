package src

// 暴力匹配
func bf(s, p string) int {
	n := len(s)
	m := len(p)
	if n < m {
		return -1
	}

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m {
			// 如果主串与模式串不匹配，则主串向右移动一个字符,模式串从头开始匹配
			if s[i+j] != p[j] {
				break
			}
			j++
		}
		if j == m {
			return i
		}
	}

	return -1
}

func kmp(s, p string) int {
	n := len(s)
	m := len(p)
	if n < m {
		return -1
	}
	next := getNext(p)

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m {
			// 暴力匹配算法当模式串和主串不匹配时，主串匹配下标 +1，模式串匹配下标置为 0，
			// KMP 算法优化点在于将模式串下标置为不匹配字符前一个字符对应 next 数组的值
			if s[i+j] != p[j] {
				// 当模式串与主串不匹配时，如果不匹配字符对应模式串下标大于 j > 0 (非首个模式串字符)，
				// 并且此字符前一个字符对应字符串部分匹配表中的值 next[j - 1] 也大于 0，
				// j - next[j - 1] 即模式串为后移的位数，等价于 j 置为 next[j - 1]

				if j != 0 && next[j-1] != 0 {
					// j 后移 j-next[j-1]，等价于 j = next[j-1]
					j = next[j-1]
				} else {
					break
				}
			}
			j++
		}
		if j == m {
			return i
		}
	}

	return -1
}

func getNext(p string) []int {
	n := len(p)
	next := make([]int, n)
	next[0] = 0
	k := 0

	// 根据已知 next 数组的前 i-1 位推测第 i 位
	for i := 1; i < n; i++ {
		for k > 0 && p[k] != p[i] {
			// k 为 p[0, i) 子串最大匹配前后缀长度
			// p[0, k) 为 p[0, i) 子串最大匹配前缀子串

			// 若：1、p[k] != p[i]，则求 p[0, i] 子串最大匹配前后缀长度问题
			// 转换成了求 p[0, k) 子串最大匹配前后缀长度问题
			// 循环直到 p[k] == p[i] (下一步处理) 或 k == 0
			k = next[k]
		}

		// 若：2、p[k] == p[i]，则 p[0, i] 子串最大匹配前后缀长度为 k + 1
		if p[k] == p[i] {
			k++
		}

		next[i] = k
	}
	return next
}
