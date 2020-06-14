package main

import (
	"fmt"
)

// DecToHexa ...
func DecToHexa(num int) string {
	var hex = make([]rune, 0, 1000)
	var flag bool

	if num < 0 {
		num = -1 * num
		flag = true
	}

	for num > 0 {
		tmp := num % 16
		if tmp < 10 {
			hex = append(hex, rune(tmp) + '0')
		} else {
			hex = append(hex, rune(tmp - 10) + 'A')
		}

		num = num / 16
	}
	
	if flag { hex = append(hex, '-') }

	// reverse hex to get correct output
	for i := 0; i < len(hex)/2; i++ {
		j := len(hex)-i-1 
		hex[i], hex[j] = hex[j], hex[i]
	}
	return string(hex)
}

func main() {
	var n, num int
	
	// Read number of input
	fmt.Println("number of input:")
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		// Read input num (decimal 10)
		fmt.Scan(&num)

		// Execute solution
		ans := DecToHexa(num)
		fmt.Println("hex: ", ans)
	}
}
