package main

import "fmt"

var res = map[[2]int]int{}

func coins(n int) int {
	return coins_helper(n, n)
}

func coins_helper(n, max int) (cnt int) {
	if n == 0 {
		return 1
	}
	if max > n {
		max = n
	}
	if cnt, ok := res[[2]int{n, max}]; ok {
		return cnt
	}
	for i := 1; i <= n; i++ {
		if i > max {
			break
		}
		cnt += coins_helper(n-i, i)
	}
	res[[2]int{n, max}] = cnt
	return cnt
}

func main() {
	n := 0
	for {
		n++
		if coins(n)%1000000 == 0 {
			fmt.Println(n)
			break
		}
	}
}
