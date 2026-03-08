package main
func CharacterConversion(str1, str2 []byte) int {

	dp := make([][]int, len(str1)+1)
	for i := range dp {
		dp[i] = make([]int, len(str2)+1)
	}

	for i := 0; i <= len(str1); i++ {
		dp[i][0] = i
	}

	for j := 0; j <= len(str2); j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {

			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {

				insert := dp[i][j-1] + 1
				delete := dp[i-1][j] + 1
				replace := dp[i-1][j-1] + 1

				dp[i][j] = min(insert, delete, replace)
			}
		}
	}

	return dp[len(str1)][len(str2)]
}

func min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}