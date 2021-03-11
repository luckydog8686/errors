package errors


type ErrorInfo string

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
	return ErrorInfo(text)
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e ErrorInfo) Error() string {
	return string(e)
}
