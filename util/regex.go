package util

import "regexp"

func IsNumberValid(n string) bool {
	expectedRegex := regexp.MustCompile(`^(?:(?:(?:\\+?|00)(98))|(0))?((?:90|91|92|93|99)[0-9]{8})$`)
	return expectedRegex.MatchString(n)
}
