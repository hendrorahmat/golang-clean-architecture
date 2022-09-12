package errors

const (
	UNKNOWN_ERROR                     ErrorCode = 0
	DATA_INVALID                      ErrorCode = 4_001_00_00001
	INVALID_HEADER_X_BUYER_ID         ErrorCode = 4_001_00_00002
	INVALID_HEADER_X_USER_ID          ErrorCode = 4_001_00_00003
	INVALID_HEADER_X_SELLER_ID        ErrorCode = 4_001_00_00004
	INVALID_REQUEST_RETRIEVE_PROVINCE ErrorCode = 4_001_00_00005
	FAILED_RETRIEVE_PROVINCE          ErrorCode = 4_001_00_00006
	FAILED_RETRIEVE_LOCATION          ErrorCode = 4_001_00_00007
	INVALID_REQUEST_RETRIEVE_LOCATION ErrorCode = 4_001_00_00008
)

var errorCodes = map[ErrorCode]*CommonError{
	UNKNOWN_ERROR: {
		ClientMessage: "Unknown error.",
		SystemMessage: "Unknown error.",
		ErrorCode:     UNKNOWN_ERROR,
	},
	DATA_INVALID: {
		ClientMessage: "Invalid Data Request",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     DATA_INVALID,
	},
	INVALID_HEADER_X_BUYER_ID: {
		ClientMessage: "Invalid buyer.",
		SystemMessage: "Invalid value of header X-Buyer-ID.",
		ErrorCode:     INVALID_HEADER_X_BUYER_ID,
	},
	INVALID_HEADER_X_USER_ID: {
		ClientMessage: "Invalid user.",
		SystemMessage: "Invalid value of header X-User-ID.",
		ErrorCode:     INVALID_HEADER_X_USER_ID,
	},
	INVALID_HEADER_X_SELLER_ID: {
		ClientMessage: "Invalid seller.",
		SystemMessage: "Invalid value of header X-Seller-ID.",
		ErrorCode:     INVALID_HEADER_X_SELLER_ID,
	},
	INVALID_REQUEST_RETRIEVE_PROVINCE: {
		ClientMessage: "Failed to request province.",
		SystemMessage: "Request has an invalid query params and/or payload to retrieve province.",
		ErrorCode:     INVALID_REQUEST_RETRIEVE_PROVINCE,
	},
	FAILED_RETRIEVE_PROVINCE: {
		ClientMessage: "Failed to retrieve province.",
		SystemMessage: "Something wrong happened while retrieve province.",
		ErrorCode:     INVALID_REQUEST_RETRIEVE_PROVINCE,
	},
	FAILED_RETRIEVE_LOCATION: {
		ClientMessage: "Failed to retrieve province.",
		SystemMessage: "Something wrong happened while retrieve province.",
		ErrorCode:     INVALID_REQUEST_RETRIEVE_LOCATION,
	},
}
