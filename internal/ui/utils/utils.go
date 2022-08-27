package utils

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TruncateString(str string, num int) string {
	truncated := str
	if num <= 3 {
		return str
	}
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		truncated = str[0:num] + "…"
	}
	return truncated
}

func TruncateStringTrailing(str string, num int) string {
	truncated := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		skipped := len(str) - num
		truncated = "..." + str[skipped:]
	}
	return truncated
}

func BoolPtr(b bool) *bool       { return &b }
func StringPtr(s string) *string { return &s }
func UintPtr(u uint) *uint       { return &u }
func IntPtr(u int) *int          { return &u }
