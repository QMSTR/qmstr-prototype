package analysis

type Analyzable interface {
	StoreResult(result map[string]interface{}) error
	GetFile() string
}

type Analyzer interface {
	GetName() string
	Analyze(aw Analyzable) error
}
