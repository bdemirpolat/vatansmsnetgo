package vatansms

import "regexp"

func PhoneVerify(phone string) string {
	re := regexp.MustCompile("[0-9]") //
	noChar := re.FindAllString(phone, -1)
	if len(noChar) < 10 {
		return ""
	}
	getLast10 := noChar[len(noChar)-10:]
	number := ""
	for char := range getLast10 {
		number += getLast10[char]
	}
	if len(number) != 10 {
		return ""
	}
	return "90" + number
}
