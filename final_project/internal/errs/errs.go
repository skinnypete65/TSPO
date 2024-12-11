package errs

import "errors"

var ErrGoodNotFound = errors.New("good not found")

var ErrUserNotExists = errors.New("user not found")
var ErrAlreadyExists = errors.New("user already exists")
var ErrInvalidPass = errors.New("invalid password")
