package stringutils

import "testing"

func Test_VerifyEmail(t *testing.T) {

	if result := VerifyEmail("contack@gmail.com"); result {
		t.Log("email is ok")
	} else {
		t.Error("Invalid email")
	}
}
func Test_VerifyEmail2(t *testing.T) {
	if result := VerifyEmail("contackidrmfly.com"); result {
		t.Log("email is ok")
	} else {
		t.Error("Invalid email")
	}
}
