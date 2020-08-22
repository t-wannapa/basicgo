package main

import (
	"errors"
	"fmt"
)

type BusinessError struct {
	Err       error
	Code      int
	Serverity int
}

func (b *BusinessError) Error() string {
	fmt.Printf("err=%s, code=%d, serverity=%d", "b.Err.ErrError()", b.Code, b.Serverity)
	return ""
}

func PrintErr(err error) {
	fmt.Println(err)
}

func main() {
	var err error = errors.New("My error msg")
	PrintErr(err)

	brr := &BusinessError{Err: err, Code: 555, Serverity: 55}
	PrintErr(brr)

}
