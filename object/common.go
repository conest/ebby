package object

// Abs : 绝对值
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Sign : 根据输入的正负情况返回 -1, 0, 1
func Sign(a int) int {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return 1
	}
	return 0
}
