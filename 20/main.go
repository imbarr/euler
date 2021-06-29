package main

func main() {
	num := 	make([]int, 1000)[:1]

	num[0] = 1
	for i := 1; i <= 100; i++ {
		carry := 0
		for j, _ := range num {
			div := num[j] * i + carry
			carry = div / 10
			num[j] = div % 10
		}
		if carry > 0 {
			num = num[:len(num) + 1]
			num[len(num) - 1] = carry % 10
			if carry / 10 > 0 {
				num = num[:len(num) + 1]
				num[len(num) - 1] = carry / 10
			}
		}
	}

	sum := 0
	for _, el := range num {
		sum += el
	}
	println(sum)
}
