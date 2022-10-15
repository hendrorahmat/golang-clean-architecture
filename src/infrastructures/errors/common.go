package errors

type ErrorCode uint

type Errorlists []string
type ValidationErrors map[string]Errorlists

type GeneralError struct {
	Message          string           `json:"message"`
	SystemMessage    string           `json:"errorMessages"`
	ValidationErrors ValidationErrors `json:"errors"`
	ErrorCode        ErrorCode        `json:"code"`
	StatusCode       int              `json:"-"`
}

func NewError(errCode ErrorCode) *GeneralError {
	var clientMessage = "Unknown error."
	var systemMessage = "Unknown error."
	var commonError = errorCodes[errCode]

	if commonError == nil {
		return &GeneralError{
			Message:       clientMessage,
			SystemMessage: systemMessage,
			ErrorCode:     errCode,
		}
	}

	return &GeneralError{
		Message:          commonError.Message,
		SystemMessage:    commonError.SystemMessage,
		ErrorCode:        errCode,
		ValidationErrors: make(ValidationErrors),
	}
}

func (err *GeneralError) SetClientMessage(message string) {
	err.Message = message
}

func (err *GeneralError) SetSystemMessage(message string) {
	err.SystemMessage = message
}

func (err *GeneralError) SetStatusCode(code int) {
	err.StatusCode = code
}
