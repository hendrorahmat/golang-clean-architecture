package errors

// urgency_service_error-type_unique-id
/**
   	urgency = p0, p1, p2, p3 -> smaller is more urgent
	service_id = 01, 02 ,...
	error_type =
		General Validation Payload Error (000),
		DB Error (001),
		Third party error (003),
		Internal Server error (004),
		/{resource_name} (005++)...
	unique_id = 001
*/

const (
	UnknownError             ErrorCode = 0
	DataPayloadInvalid       ErrorCode = 4_01_000_001
	QueryParamDataInvalid    ErrorCode = 4_01_000_002
	InvalidHeaderXDeviceType ErrorCode = 4_01_000_003
)

var errorCodes = map[ErrorCode]*GeneralError{
	UnknownError: {
		Message:       "Unknown error.",
		SystemMessage: "Unknown error.",
		ErrorCode:     UnknownError,
	},
	QueryParamDataInvalid: {
		Message:       "Invalid Data Request",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     QueryParamDataInvalid,
	},
	InvalidHeaderXDeviceType: {
		Message:       "Invalid header.",
		SystemMessage: "Invalid value of header X-Device-Type.",
		ErrorCode:     InvalidHeaderXDeviceType,
	},
	DataPayloadInvalid: {
		Message:       "Invalid Data Request",
		SystemMessage: "Some of data payload has invalid value.",
		ErrorCode:     DataPayloadInvalid,
	},
}
