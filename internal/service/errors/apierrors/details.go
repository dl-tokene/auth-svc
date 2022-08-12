package apierrors

// detailsMap provides implicit mapping of `ErrorObject.Code` to `ErrorObject.Detail`
// This is just for convenience so we don't have to type out wordy descriptions in handlers or wherever else
// It is not required to fill in details for every code
var detailsMap = map[string]string{
	// Handlers
	CodeEmptyBody:      "Request body was expected.",
	CodeBadRequestData: "Some data in the request was invalid. Please refer to the service documentation.",
	// Login
	CodeNotRegistered: "User with such credentials is not registered. Please register first.",
	// Email Verification
	CodeInvalidEmailVerification: "The verification code didn't match the email. Please check if the email is correct.",
	// Quotes
	CodeFailedPriceLookup: "Failed to look up exchange prices. Either the exchange API is unavailable or your request is invalid.",
	// Address Deletion
	CodeActiveAddressDeactivation: "It is not possible to unlink the address you're currently logged in with. Please log in with another address or contact support team,",
	// Yoti
	CodeFailedYotiExchange: "Failed to exchange Yoti access token for user's personal details. Invalid access token?",
	CodeFailedYotiVerification: "User's personal information didn't match service's requirements. See `Meta` for more information.",
	// Signature Verification
	CodeNonceNotFound:               "No nonce was found for you address. It either has expired or was never requested",
	CodeSignatureVerificationFailed: "Signature Verification failed. The signature was invalid or addresses didn't match.",
	// Session Authentication
	CodeSessionTokenNotFound: "Session token wasn't found. Is Authorization header correct? Are cookies enabled?",
	CodeSessionTokenInvalid:  "Session token is invalid. It either has expired or is corrupted. Please log in and obtain a new one.",
	// Resource Dependency
	CodeResourceNotFound: "The resource required to perform your operation was not found.",
	// Conflicts
	CodeAddressExists: "Provided address is already associated with one of the existing accounts. Please use another one or contact system adminstrators.",
	// Suspended/Rejected User
	CodeUserSuspended: "This account has been suspended and cannot perform any actions on the platform at the moment.",
	CodeUserRejected: "This account has been rejected from the platform and lacks necessary permissions to perform the request or action",
	// Verification Deletion
	CodeVerificationDeletionNotSupported: "Deletion of verifications with provided type and status is not currently supported. Please refer to the documentation.",
	// Faucets
	CodeFaucetAlreadyUsed: "You've already used this faucet.",
	// Not owned resources
	CodeResourceNotOwned: "Resource is not owned by this account.",
	CodeNotAddressOwner: "The specified address doesn't belong to the authenticated account.",
	// General Errors
	CodeBadRequest:    "The request was invalid in some way.",
	CodeUnauthorized:  "You're not authorized. Please register or log in.",
	CodeForbidden:     "The requester lacks necessary permissions to perform the request or action.",
	CodeNotFound:      "The requested resource could not be found.",
	CodeConflict:      "The resource you were trying to create either already exists or is known to the service.",
	CodeTooEarly:      "You're rushing things. Don't.",
	CodeInternalError: "Some internal error occurred. Please report this error to service maintainers.",
}
