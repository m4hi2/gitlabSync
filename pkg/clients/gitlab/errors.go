package gitlab

import (
	"github.com/m4hi2/gitlabSync/pkg/utils/glserrors"
	"net/http"
)

var (
	ErrUnAuthorized        = &glserrors.GLSError{Code: "GLS401", Message: "Unauthorized Gitlab User", Err: nil}
	ErrForbidden           = &glserrors.GLSError{Code: "GLS403", Message: "Forbidden Request", Err: nil}
	ErrGroupNotFound       = &glserrors.GLSError{Code: "GLS404", Message: "Gitlab Group not found", Err: nil}
	ErrInternalServerError = &glserrors.GLSError{Code: "GLS500", Message: "Internal Server Error in Gitlab.", Err: nil}
	ErrCallingGitlab       = &glserrors.GLSError{Code: "GLS600", Message: "Client Can not connect to Gitlab.", Err: nil}
	ErrReadingBody         = &glserrors.GLSError{Code: "GLS601", Message: "Gitlab Reading Body Error", Err: nil}
	ErrUnmarshalError      = &glserrors.GLSError{Code: "GLS602", Message: "Can not unmarshal Gitlab data", Err: nil}
)

func httpResponseErrorParse(respCode int) *glserrors.GLSError {
	if respCode == http.StatusUnauthorized {
		return ErrUnAuthorized
	}

	if respCode == http.StatusForbidden {
		return ErrForbidden
	}

	if respCode == http.StatusInternalServerError {
		return ErrInternalServerError
	}

	if respCode == http.StatusNotFound {
		return ErrGroupNotFound
	}

	return nil
}
