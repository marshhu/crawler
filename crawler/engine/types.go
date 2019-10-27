package engine

//Request struct
type Request struct{
	URL string 
	ParserFunc func([]byte) ParseResult
}

//ParseResult struct
type ParseResult struct {
	Requests []Request
	Items []interface{}
}
//NilParser fucntion
func NilParser([]byte) ParseResult{
	return ParseResult{}
}