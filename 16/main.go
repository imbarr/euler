package main

func main() {
	num := 	make([]byte, 1000)[:1]

	num[0] = 1
	for i := 0; i < 1000; i++ {
		var carry byte = 0
		for j, _ := range num {
			div := num[j] * 2 + carry
			carry = div / 10
			num[j] = div % 10
		}
		if carry > 0 {
			num = num[:len(num) + 1]
			num[len(num) - 1] = carry
		}
	}

	sum := 0
	for _, el := range num {
		sum += int(el)
	}
	println(sum)
}
