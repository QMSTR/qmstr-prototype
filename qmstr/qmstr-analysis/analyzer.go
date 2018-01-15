package analysis

type Analyzable interface {
	StoreResult(result map[string]interface{}) error
	GetFile() string
}

type Analyzer interface {
	Configure(data map[string]interface{}) error
	GetName() string
	Analyze(aw Analyzable) error
}
