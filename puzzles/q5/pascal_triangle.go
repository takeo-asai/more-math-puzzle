package q5

func PascalTriangle(n int) []int {
	if n == 0 {
		return []int{1}
	}
	if n == 1 {
		return []int{1, 1}
	}
	previousOne := PascalTriangle(n - 1)
	nextOne := make([]int, len(previousOne)+1)
	nextOne[0] = 1
	for i := 0; i < len(previousOne)-1; i++ {
		nextOne[i+1] = previousOne[i] + previousOne[i+1]
	}
	nextOne[len(previousOne)] = 1
	return nextOne
}
