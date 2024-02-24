package helper

import "errors"

func HttpStatusErr(err error) int {
	switch {
	case errors.Is(err, ErrRegisterUser):
		return 401
	default:
		return 500
	}
}
