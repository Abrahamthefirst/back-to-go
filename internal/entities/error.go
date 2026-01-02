package entities

import "errors"

var (
	RecordNotFound        = errors.New("Record Not Found")
	ErrConflict           = errors.New("")
	BadRequestException   = errors.New("Bad request")
	ErrInvalidCredentials = errors.New("Invalid Credentials")
	ErrInternal           = errors.New("Internal Server Error")
)
