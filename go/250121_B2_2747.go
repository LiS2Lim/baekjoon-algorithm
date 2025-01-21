/**
 * https://www.acmicpc.net/problem/2747
 */
package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)
	
	dp := make([]int, n+1)
	if n >= 1 {
		dp[1] = 1;
	}
	
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	
	fmt.Printf("%d\n", dp[n])
}
