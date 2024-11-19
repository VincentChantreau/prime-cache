package parser

type ParserConfig struct {
	FilteredFileExtensions []string
}
type Parser struct {
	Config *ParserConfig
}
