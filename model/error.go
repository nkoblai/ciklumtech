package model

// HTTPError represents custom HTTP error type
type HTTPError struct {
	ErrorMessage string
}

// NewHTTPError creates custom HTTP error type
func NewHTTPError(msg string) error {
	if msg == "" {
		return nil
	}
	return &HTTPError{msg}
}

// Error returns error as a string
func (e HTTPError) Error() string {
	return e.ErrorMessage
}
