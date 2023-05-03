package error

const (
	TypeAfter         = "AFTER"
	TypeAfterOrEqual  = "AFTER_OR_EQUAL"
	TypeBefore        = "BEFORE"
	TypeBeforeOrEqual = "BEFORE_OR_EQUAL"
	TypeBoolean       = "BOOLEAN"
	TypeCustom        = "CUSTOM"
	TypeDateFormat    = "DATE_FORMAT"
	TypeDuration      = "DURATION"
	TypeEmail         = "EMAIL"
	TypeIn            = "IN"
	TypeInt           = "INT"
	TypeLength        = "LENGTH"
	TypeMax           = "MAX"
	TypeMin           = "MIN"
	TypeRequired      = "REQUIRED"
	TypeSlice         = "ARRAY"
	TypeSliceOf       = "ARRAY_OF"
	TypeString        = "STRING"
)

const (
	SubtypeNumber = "NUMBER"
	SubtypeString = "STRING"
	SubtypeSlice  = "ARRAY"
	SubtypeMap    = "MAP"
)
