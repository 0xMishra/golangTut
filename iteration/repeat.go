package iteration

func Repeat(s string, n int) string {
	sol := ""
	for i := 0; i < n; i++ {
		sol += s
	}
	return sol
}
