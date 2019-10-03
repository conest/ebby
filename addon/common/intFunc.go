package common

// Abs : 绝对值
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Clamp : 范围限制
func Clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
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
