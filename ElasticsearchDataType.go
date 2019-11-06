package elasticsearch

//EsType is elasticsearch Types
type esType string

//EsType list
const (
	esNull        esType = "null"
	esBoolean     esType = "boolean"
	esByte        esType = "byte"
	esShort       esType = "short"
	esInteger     esType = "integer"
	esLong        esType = "long"
	esDouble      esType = "double"
	esFloat       esType = "float"
	esHalfFloat   esType = "half_float"
	esScaledFloat esType = "scaled_float"
	esKeyword     esType = "keyword"
	esText        esType = "text"
	esBinary      esType = "binary"
	esDatetime    esType = "datetime"
	esIP          esType = "ip"
	esObject      esType = "object"
	esNested      esType = "nested"
	esUnsupported esType = "unsupported"
)
