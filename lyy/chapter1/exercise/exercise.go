package chapter1

import "errors"

func Cheng(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("约定乘数或被乘数不能为0")
	}
	return a * b, nil
}
