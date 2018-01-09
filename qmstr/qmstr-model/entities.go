package model

/* The package model contains data model entities that are shared
/* between the master and the different client programs. */

// Entity is the base type that can be stored in the data model.
type Entity interface {
	// Id identifies the entitiy. It needs to be unique for each type of entity.
	ID() string
}

// SourceEntity represents a source code file that becomes part of a target.
type SourceEntity struct {
	Path     string
	Hash     string
	Licenses []string
}

// ID for SourceEntity uses the path (assuming it is the package local
// path of the source file).
func (e *SourceEntity) ID() string {
	return e.Path
}

func (e *SourceEntity) StoreResult(result map[string]interface{}) error {
	if lics, ok := result["licenses"]; ok {
		e.Licenses = lics.([]string)
	}
	return nil
}

func (e *SourceEntity) GetFile() string {
	return e.Path
}

// DependencyEntity represents a dependency for a target.
type DependencyEntity struct {
	Name string
	Hash string
}

// ID for DependencyEntity uses the name (for the moment).
func (e *DependencyEntity) ID() string { return e.Name }

// TargetEntity represents a target that is generated during the software build process.
type TargetEntity struct {
	Name         string
	Hash         string
	Sources      []string
	Dependencies []string
	Linked       bool
	Path         string
}

// ID for TargetEntity uses the name (for the moment).
func (e *TargetEntity) ID() string { return e.Name }
