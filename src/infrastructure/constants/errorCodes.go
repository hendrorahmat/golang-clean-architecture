package constants

// urgency_service_error-type_unique-id
/**
   	urgency = p0, p1, p2, p3 -> smaller is more urgent
	service_id = 01, 02 ,...
	error_type =
		General Validation Payload / business logic Error (001),
		DBGorm Error (002),
		Third party error (003),
		Internal Server error (004),
		{resource_name} = (005++)...
	unique_id = 001
*/

/* Category Infrastructures */
const (
	DBQueryNotFoundCode uint = 3_01_001_004
	DBQueryErrorCode    uint = 0_01_002_001
	InternalCodeError   uint = 0_01_004_001
)

/* Category Request Param */
const (
	InvalidHeaderXDeviceTypeCode uint = 3_01_001_003
	InvalidHeaderCode            uint = 3_01_001_003
	QueryParamDataInvalidCode    uint = 3_01_001_002
	FieldMissingCode             uint = 3_01_001_005
)

/* Category Business Logic */
const (
	DataPayloadInvalidCode uint = 3_01_001_001
	FieldsRequiredCode     uint = 3_01_001_002
)
