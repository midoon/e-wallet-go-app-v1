package helper

import "errors"

var ErrRegisterUser = errors.New("error register failed")
var ErrDuplicateData = errors.New("error data already exist")
