package util

func ToInt(input string) (res int) {
	mult := 1
	for i := len(input) - 1; i >= 0; i-- {
		res += int(input[i]-'0') * mult
		mult *= 10
	}
	return res
}
