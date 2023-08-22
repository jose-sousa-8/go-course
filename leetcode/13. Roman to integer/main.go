package main

import "fmt"

func romanToInt(s string) int {
	var symbols = make(map[byte]int)
	symbols['I'] = 1
	symbols['V'] = 5
	symbols['X'] = 10
	symbols['L'] = 50
	symbols['C'] = 100
	symbols['D'] = 500
	symbols['M'] = 1000

	var total int
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			total += symbols[s[i]]
			continue
		}
		if s[i] == 'I' {
			if s[i+1] == 'X' {
				total += 9
				i++
				continue
			}
			if s[i+1] == 'V' {
				total += 4
				i++
				continue
			}
			total += 1
			continue
		}
		if s[i] == 'X' {
			if s[i+1] == 'L' {
				total += 40
				i++
				continue
			}
			if s[i+1] == 'C' {
				total += 90
				i++
				continue
			}
			total += 10
			continue
		}
		if s[i] == 'C' {
			if s[i+1] == 'D' {
				total += 400
				i++
				continue
			}
			if s[i+1] == 'M' {
				total += 900
				i++
				continue
			}
			total += 100
			continue
		}

		total += symbols[s[i]]
	}

	return total
}

func main() {
	fmt.Println(romanToInt("MCMXCIII"))
}
