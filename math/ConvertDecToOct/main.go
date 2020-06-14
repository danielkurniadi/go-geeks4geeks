package main

import (
	"fmt"
)

func DecToOct(num int) string {
	var oct = make([]rune, 0, 1000)
	var flag bool

	if num == 0 { return "0" }

	if num < 0 {
		flag = true
		num = -1 * num
	}

	for num > 0 {
		tmp := num % 8
		oct = append(oct, rune(tmp) + '0')
		num /= 8
	}

	if flag { oct = append(oct, '-') }

	for i := 0; i < len(oct)/2; i++ {
		j := len(oct)-i-1
		oct[i], oct[j] = oct[j], oct[i]
	}

	return string(oct)
}

func main() {
	var n, num int
	
	// Read number of input
	fmt.Println("number of input:")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		// Read input num (decimal 10)
		fmt.Scan(&num)

		// Execute solution
		ans := DecToOct(num)
		fmt.Println("oct: ", ans)
	}
}
