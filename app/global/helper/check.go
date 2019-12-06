package helper

import (
	"github.com/pkg/errors"
	"goformat-v2/app/global/errorcode"
	"runtime/debug"
	"strconv"
)

func Atoi() (int, errorcode.Error) {
	i, err := strconv.Atoi("wer")
	if err != nil {
		return 0, errorcode.ErrorHandler("PERMISSION_DENIED", err) // annotate errors with stacktrace
	}
	return i, nil
}

func GetString(data interface{}) error {
	switch data.(type) {
	case int:
		err := errors.New("is int")
		apiError := errorcode.WrapError(err)
		return apiError
	case string:
		err := errors.New("is string")
		apiErr := errorcode.WrapError(err)
		debug.PrintStack()
		return apiErr
	}
	err := errors.New("no data")
	apiError := errorcode.WrapError(err)
	return apiError
}
