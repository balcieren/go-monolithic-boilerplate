package fail

import (
	"fmt"
	"strconv"
	"strings"
)

type Fail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(code int, message ...string) error {
	f := Fail{
		Code: code,
	}

	if len(message) > 0 {
		f.Message = message[0]
	}

	return fmt.Errorf("%d: %s", f.Code, f.Message)
}

func Convert(err error) (int, string) {
	var f Fail
	if err != nil {
		text := strings.Split(err.Error(), ": ")
		code, _ := strconv.Atoi(text[0])
		f.Code = code
		f.Message = text[1]
	}

	return f.Code, f.Message
}
