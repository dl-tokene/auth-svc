package errors

type idpCheckErrorKey int

// These errors are based on official yes® Relaying Party Developer Guide
// https://www.yes.com/docs/rp-devguide/2.0/IDENTITY/index.html#_issuer_uri_check
const (
	// HTTP 204 -> IDP found and active
	InvalidIssuerURL idpCheckErrorKey = iota // HTTP 400
	IDPNotFound                              // HTTP 404
	InactiveIDP                              // HTTP 423
	UnknownError                             // Anything else
)

type IDPCheckError struct {
	IssuerURL string
	Type      idpCheckErrorKey
}

func (e IDPCheckError) Error() string {
	switch e.Type {
	case InvalidIssuerURL:
		return "invalid issuer URL: " + e.IssuerURL
	case IDPNotFound:
		return "IDP not found for issuer URL: " + e.IssuerURL
	case InactiveIDP:
		return "IDP is inactive for issuer URL: " + e.IssuerURL
	case UnknownError:
		return "yes.com returned unkown status code for issuer URL: " + e.IssuerURL
	}
	return "unhandled error type"
}

func NewIDPError(issuerUrl string, statusCode int) error {
	var errorType idpCheckErrorKey
	switch statusCode {
	case 400:
		errorType = InvalidIssuerURL
	case 404:
		errorType = IDPNotFound
	case 423:
		errorType = InactiveIDP
	default:
		errorType = UnknownError
	}
	return IDPCheckError{IssuerURL: issuerUrl, Type: errorType}
}
