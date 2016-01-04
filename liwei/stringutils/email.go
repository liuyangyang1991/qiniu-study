package stringutils

import (
	"regexp"
)

func VerifyEmail(emailAddress string) bool {
	reg := regexp.MustCompile(`^(\w)+([\.|-]\w+)*@(\w)+((\.\w+)+)$`)
	return reg.MatchString(emailAddress)
}
