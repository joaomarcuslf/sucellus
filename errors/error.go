package errors

import (
	"fmt"
)

var ErrorEmum = map[string]string{
	"MONGO_ERROR":      "error: Mongo Error",
	"MIGRATION_ERROR":  "error: Migration could not be completed",
	"REPOSITORY_ERROR": "error: Error in repository execution",
	"VALIDATION_ERROR": "error: Error in some input",
}

type SystemError struct {
	Type    string
	Message string
	Err     error
}

func (e *SystemError) Error() string {
	return fmt.Sprintf(
		"%s\n  Message: %s\n  Err:\n    %s",
		ErrorEmum[e.Type],
		e.Message,
		e.Err.Error(),
	)
}

func FormatError(errorId, message string, additionalError error) error {
	return &SystemError{
		Type:    errorId,
		Message: message,
		Err:     additionalError,
	}
}
