package excep

// FORMAT: GENERAL_EXCEPTION_CODE or
// FORMAT: GENERAL_EXCEPTION_CODE.DOMAIN.DOMAIN_SPECIFIC_CODE or
// FORMAT: DOMAIN.DOMAIN_SPECIFIC_CODE
const (
	EXCEP_UNAUTHORIZED             string = "UNAUTHORIZED"
	EXCEP_INVALID_PAYLOAD          string = "INVALID_PAYLOAD"
	EXCEP_INTERNAL_SERVER_ERROR    string = "INTERNAL_SERVER_ERROR"

	EXCEP_USER_EMAIL_ALREADY_TAKEN string = "USER.EMAIL_ALREADY_TAKEN"
)

type DomainExcep struct {
	Type    string `json:"type"`
	Details string `json:"details"`
}

type HttpExcep struct {
	Status  int    `json:"status"`
	Type    string `json:"type"`
	Details string `json:"details"`
}

func (d DomainExcep) Error() string {
    return d.Details
}

func NewAuthExcep(details string) DomainExcep {
    if details == "" {
        details = "unauthorized"
    }

    return DomainExcep{
        Type: EXCEP_UNAUTHORIZED,
        Details: details,
    }
}

func NewInternalExcep(details string) DomainExcep {
    if details == "" {
        details = "internal server exception"
    }

    return DomainExcep{
        Type: EXCEP_UNAUTHORIZED,
        Details: details,
    }
}

func NewInvalidExcep(details string) DomainExcep {
    if details == "" {
        details = "invalid payload"
    }

    return DomainExcep{
        Type: EXCEP_INVALID_PAYLOAD,
        Details: details,
    }
}
