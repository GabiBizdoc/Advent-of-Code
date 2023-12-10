package env

type AppConfigError struct {
	Message string
}

func NewAppConfigError(message string) *AppConfigError {
	return &AppConfigError{Message: message}
}

func (e *AppConfigError) Error() string {
	return e.Message
}
