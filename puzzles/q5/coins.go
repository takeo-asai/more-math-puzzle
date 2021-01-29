package q5

func Coins(n int) int {
	if n < 0 {
		return 0
	}
	if n >= 10000 {
		return n/10000 + Coins(n%10000)
	}
	if n >= 5000 {
		return 1 + Coins(n-5000)
	}
	if n >= 2000 {
		return 1 + Coins(n-2000)
	}
	if n >= 1000 {
		return 1 + Coins(n-1000)
	}
	if n >= 500 {
		return 1 + Coins(n-500)
	}
	if n >= 100 {
		return 1 + Coins(n-100)
	}
	if n >= 50 {
		return 1 + Coins(n-50)
	}
	if n >= 10 {
		return 1 + Coins(n-10)
	}
	if n >= 5 {
		return 1 + Coins(n-5)
	}
	return n
}
