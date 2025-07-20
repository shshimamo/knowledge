package util

import "errors"

var ErrUnauthorized = errors.New("unauthorized")

var ErrForbidden = errors.New("forbidden")