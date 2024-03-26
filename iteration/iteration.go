package iteration

func Repeat(str string) string {
	res := ""
	for range 5 {
		res += str
	}
	return res
}
