package i18n

type MessageID string

const (
	DefaultError                      string    = "Error"
	DefaultSuccessful                 string    = "Successful"
	ErrServerValidatorNotAvailable    MessageID = "ERR_SERVER_VALIDATOR_NOT_AVAILABLE"
	ErrServerReadingRequestBody       MessageID = "ERR_SERVER_READING_REQUEST_BODY"
	ErrRequestParseBody               MessageID = "ERR_REQUEST_PARSE_BODY"
	ErrRequestValidatorGeneral        MessageID = "ERR_REQUEST_VALIDATOR_GENERAL"
	ErrRequestValidatorDefault        MessageID = "ERR_REQUEST_VALIDATOR_DEFAULT"
	ErrRequestValidatorAlphanum       MessageID = "ERR_REQUEST_VALIDATOR_ALPHANUM"
	ErrRequestValidatorMin            MessageID = "ERR_REQUEST_VALIDATOR_MIN"
	ErrRequestValidatorMax            MessageID = "ERR_REQUEST_VALIDATOR_MAX"
	ErrRequestValidatorSecurepassword MessageID = "ERR_REQUEST_VALIDATOR_SECUREPASSWORD"
	ErrUserCreateRegister             MessageID = "ERR_USER_CREATE_REGISTER"
)
