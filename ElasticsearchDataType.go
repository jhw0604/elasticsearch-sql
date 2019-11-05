package elasticsearch

type elasticsearchType string

//elasticsearchType is elasticsearch Types
const (
	ElaTypeNull        elasticsearchType = "null"
	ElaTypeBoolean     elasticsearchType = "boolean"
	ElaTypeByte        elasticsearchType = "byte"
	ElaTypeShort       elasticsearchType = "short"
	ElaTypeInteger     elasticsearchType = "integer"
	ElaTypeLong        elasticsearchType = "long"
	ElaTypeDouble      elasticsearchType = "double"
	ElaTypeFloat       elasticsearchType = "float"
	ElaTypeHalfFloat   elasticsearchType = "half_float"
	ElaTypeScaledFloat elasticsearchType = "scaled_float"
	ElaTypeKeyword     elasticsearchType = "keyword"
	ElaTypeText        elasticsearchType = "text"
	ElaTypeBinary      elasticsearchType = "binary"
	ElaTypeDatetime    elasticsearchType = "datetime"
	ElaTypeIP          elasticsearchType = "ip"
	ElaTypeObject      elasticsearchType = "object"
	ElaTypeNested      elasticsearchType = "nested"
	ElaTypeUnsupported elasticsearchType = "unsupported"
)
