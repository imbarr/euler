package main

func main() {
	for i := 20; ; i++ {
		flag := true
		for d := 2; d <= 20 && flag; d++ {
			if i % d != 0 {
				flag = false
			}
		}
		if flag {
			println(i)
			return
		}
	}
}
