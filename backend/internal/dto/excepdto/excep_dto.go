package excepdto

import (
	"net/http"
	"sharedlambdacode/internal/excep"
)

func ErrorToHttpExcep(err error) excep.HttpExcep {
	if exceptionIns, ok := err.(excep.DomainExcep); ok {
		return excep.HttpExcep{
			Status:  getHttpStatusForExceptionType(exceptionIns.Type),
			Type:    exceptionIns.Type,
			Details: exceptionIns.Details,
		}
	}

	return excep.HttpExcep{
		Status:  getHttpStatusForExceptionType(""),
		Type:    excep.EXCEP_INTERNAL_SERVER_ERROR,
		Details: err.Error(),
	}
}

func getHttpStatusForExceptionType(exceptionType string) int {
	switch exceptionType {
	case excep.EXCEP_INVALID_PAYLOAD,
		excep.EXCEP_USER_EMAIL_ALREADY_TAKEN:
		return http.StatusBadRequest
	case excep.EXCEP_UNAUTHORIZED:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
