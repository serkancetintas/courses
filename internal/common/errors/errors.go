package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown        = ErrorType{t: "unknown"}
	ErrorTypeAuthorization  = ErrorType{t: "authorization"}
	ErrorTypeIncorrectInput = ErrorType{t: "incorrect-input"}
)
