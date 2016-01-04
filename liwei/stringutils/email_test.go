package stringutils
import "testing"

func Test_VerifyEmail(t *testing.T) {

	if result := VerifyEmail("contack@idrmfly.com"); !result {
		t.Log("email is ok")
	}else {
		t.Error("Invalid email")
	}
}
