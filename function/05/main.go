package main

func countDigits(num int) (count int) {
	for num != 0 {
		num /= 10
		count++
	}
	return
}

func main() {
	countDigits(12345)
}
