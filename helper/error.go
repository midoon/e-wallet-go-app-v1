package helper

import "errors"

var ErrDuplicateData = errors.New("error data already exist")
var ErrValidation = errors.New("error validation")
var ErrEmailOrPaswordWrong = errors.New("error email or passowrd wrong")
