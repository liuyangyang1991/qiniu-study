package stringutils
import "regexp"


func VerifyEmail(emailAddress string) bool {
	if existed, err := regexp.MatchString(`^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$ `,
		emailAddress); !existed || err != nil {
		return false
	}
	return true
}

