package models

type HTTPError struct {
	ErrorMessage string
}

func NewHTTPError(msg string) error {
	if msg == "" {
		return nil
	}
	return &HTTPError{msg}
}

func (e HTTPError) Error() string {
	return e.ErrorMessage
}
