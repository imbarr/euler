package main

import "strings"

func toText(n int) string {
	if n / 10 == 0 {
		return toTextSingleDigit(n)
	} else if n / 100 == 0 {
		return toTextDoubleDigit(n)
	} else {
		return toTextTripleDigit(n)
	}
}

func toTextTripleDigit(n int) string {
	start := toTextSingleDigit(n / 100) + " hundred"
	if n % 100 == 0 {
		return start
	} else {
		return start + " and " + toTextDoubleDigit(n % 100)
	}
}

func toTextDoubleDigit(n int) string {
	if n < 10 {
		return toTextSingleDigit(n)
	} else if n < 20 {
		return toTextTeen(n)
	} else if n % 10 == 0 {
		return toTextDecimal(n / 10)
	} else {
		return toTextDecimal(n / 10) + " " + toTextSingleDigit(n % 10)
	}
}

func toTextSingleDigit(n int) string {
	switch n {
	case 1: return "one"
	case 2: return "two"
	case 3: return "three"
	case 4: return "four"
	case 5: return "five"
	case 6: return "six"
	case 7: return "seven"
	case 8: return "eight"
	case 9: return "nine"
	default: panic("")
	}
}

func toTextTeen(n int) string {
	switch n {
	case 10: return "ten"
	case 11: return "eleven"
	case 12: return "twelve"
	case 13: return "thirteen"
	case 14: return "fourteen"
	case 15: return "fifteen"
	case 16: return "sixteen"
	case 17: return "seventeen"
	case 18: return "eighteen"
	case 19: return "nineteen"
	default: panic("")
	}
}

func toTextDecimal(n int) string {
	switch n {
	case 2: return "twenty"
	case 3: return "thirty"
	case 4: return "forty"
	case 5: return "fifty"
	case 6: return "sixty"
	case 7: return "seventy"
	case 8: return "eighty"
	case 9: return "ninety"
	default: panic("")
	}
}

func main() {
	sum := 0
	for n := 1; n < 1000; n++ {
		text := toText(n)
		println(text)
		sum += len(strings.ReplaceAll(text, " ", ""))
	}
	sum += len(strings.ReplaceAll("one thousand", " ", ""))

	println(sum)
}
