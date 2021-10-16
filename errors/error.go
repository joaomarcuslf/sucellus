package errors

import "fmt"

var ErrorEmum = map[string]string{
	"MONGO_ERROR":      "error: Mongo Error",
	"MIGRATION_ERROR":  "error: Migration could not be completed",
	"REPOSITORY_ERROR": "error: Error in repository execution",
}

func FormatError(errorId, message string, additionalError error) error {
	return fmt.Errorf(
		"%s\n  Message: %s\n  Err:\n    %s",
		ErrorEmum[errorId],
		message,
		additionalError.Error(),
	)
}
