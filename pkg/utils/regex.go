package utils

import (
	"log"
	"regexp"
)

// RegexpPassword ...
func RegexpPassword(fullName string) (match bool) {
	reg, _ := regexp.Compile(`^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	log.Println(fullName)
	if !reg.MatchString(fullName) {
		return
	}

	return true
}
