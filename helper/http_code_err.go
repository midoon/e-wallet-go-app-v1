package helper

import "errors"

func HttpStatusErr(err error) int {
	switch {
	case errors.Is(err, ErrRegisterUser):
		return 401
	case errors.Is(err, ErrValidation):
		return 403
	default:
		return 500
	}
}
