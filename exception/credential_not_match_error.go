package exception

type CredentialNotMatchError struct {
	Error string
}

func NewCredentialNotMatchError(error string) CredentialNotMatchError {
	return CredentialNotMatchError{Error: error}
}
