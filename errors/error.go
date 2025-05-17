package errors

import "errors"

var (
	ErrEmaiAlreadyTaken         = errors.New("email already taken")
	ErrSomethingWentWrong       = errors.New("Something went wrong")
	ErrInvalidCredentials       = errors.New("Invalid Credentials")
	ErrEmailNotVerified         = errors.New("Email Not Verified")
	ErrCantSendVerificationMail = errors.New("Cant Resend Verification Mail")
	ErrCodeAlreadyExist         = errors.New("Change Custom Code, It Already Exists")
	ErrUrlAlreadyExist          = errors.New("Url Already Exist")
	ErrLinkExpired              = errors.New("The Link Have Already Expired")
)
