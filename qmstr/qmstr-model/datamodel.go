package model

import "fmt"

/* DataModel implements a model of entities that are stored in
/* different buckets. Relations between the entities are resolved
/* programmatically for the moment. Long-term, the implementation
/* behind the data model will be a graph database. */

// DataModel implements the project structure model from sources, dependencies and targets.
type DataModel struct {
	sources      map[string]SourceEntity
	dependencies map[string]DependencyEntity
	targets      map[string]TargetEntity
}

// NewModel creates a default-constructed DataModel.
func NewModel() *DataModel {
	m := new(DataModel)
	m.sources = make(map[string]SourceEntity)
	m.dependencies = make(map[string]DependencyEntity)
	m.targets = make(map[string]TargetEntity)
	return m
}

// GetSourceEntity retrieves a source entity from the model.
func (m *DataModel) GetSourceEntity(id string) (SourceEntity, error) {
	if value, ok := m.sources[id]; ok {
		return value, nil
	}
	return SourceEntity{"", "", []string{}}, fmt.Errorf("source entity %s does not exist", id)
}

// AddSourceEntity adds a source file to the model.
func (m *DataModel) AddSourceEntity(source SourceEntity) error {
	if _, ok := m.sources[source.ID()]; ok {
		return fmt.Errorf("source entity %s already exists", source.ID())
	}
	m.sources[source.ID()] = source
	return nil
}

// ModifySourceEntity updates an existing source entity.
func (m *DataModel) ModifySourceEntity(source SourceEntity) error {
	if _, ok := m.sources[source.ID()]; ok {
		m.sources[source.ID()] = source
		return nil
	}
	return fmt.Errorf("source entity %s not found", source.ID())
}

// DeleteSourceEntity deletes an existing source entity.
func (m *DataModel) DeleteSourceEntity(source SourceEntity) error {
	if _, ok := m.sources[source.ID()]; ok {
		delete(m.sources, source.ID())
		return nil
	}
	return fmt.Errorf("source entity %s does not exist", source.ID())
}

// GetDependencyEntity retrieves a dependency entity from the model.
func (m *DataModel) GetDependencyEntity(id string) (DependencyEntity, error) {
	if value, ok := m.dependencies[id]; ok {
		return value, nil
	}
	return DependencyEntity{"", ""}, fmt.Errorf("dependency entity %s does not exist", id)
}

// AddDependencyEntity adds a dependency file to the model.
func (m *DataModel) AddDependencyEntity(dep DependencyEntity) error {
	if _, ok := m.dependencies[dep.ID()]; ok {
		return fmt.Errorf("dependency entity %s already exists", dep.ID())
	}
	m.dependencies[dep.ID()] = dep
	return nil
}

// ModifyDependencyEntity updates an existing dependency entity.
func (m *DataModel) ModifyDependencyEntity(dep DependencyEntity) error {
	if _, ok := m.dependencies[dep.ID()]; ok {
		m.dependencies[dep.ID()] = dep
		return nil
	}
	return fmt.Errorf("dependency entity %s not found", dep.ID())
}

// DeleteDependencyEntity deletes an existing dependency entity.
func (m *DataModel) DeleteDependencyEntity(dep DependencyEntity) error {
	if _, ok := m.dependencies[dep.ID()]; ok {
		delete(m.dependencies, dep.ID())
		return nil
	}
	return fmt.Errorf("dependency entity %s does not exist", dep.ID())
}

// GetTargetEntity retrieves a target entity from the model.
func (m *DataModel) GetTargetEntity(id string) (TargetEntity, error) {
	if value, ok := m.targets[id]; ok {
		return value, nil
	}
	return TargetEntity{"", "", []string{}, []string{}, false}, fmt.Errorf("target entity %s does not exist", id)
}

// GetAllTargetEntities retrieves all target entities from the model.
func (m *DataModel) GetAllTargetEntities() []TargetEntity {
	targets := []TargetEntity{}
	for _, target := range m.targets {
		targets = append(targets, target)
	}
	return targets
}

// GetAllSourceEntities retrieves all target entities from the model.
func (m *DataModel) GetAllSourceEntities() []SourceEntity {
	srcs := []SourceEntity{}
	for _, src := range m.sources {
		srcs = append(srcs, src)
	}
	return srcs
}

// GetAllDependencyEntities retrieves all dependency entities from the model.
func (m *DataModel) GetAllDependencyEntities() []DependencyEntity {
	deps := []DependencyEntity{}
	for _, dep := range m.dependencies {
		deps = append(deps, dep)
	}
	return deps
}

// GetAllLinkedTargets retrieves all target entity from the model.
func (m *DataModel) GetAllLinkedTargets() []string {
	targets := []string{}
	for _, target := range m.targets {
		if target.Linked {
			targets = append(targets, target.Name)
		}
	}
	return targets
}

// AddTargetEntity adds a target file to the model.
func (m *DataModel) AddTargetEntity(target TargetEntity) error {
	if _, ok := m.targets[target.ID()]; ok {
		return fmt.Errorf("target entity %s already exists", target.ID())
	}
	m.targets[target.ID()] = target
	return nil
}

// ModifyTargetEntity updates an existing target entity.
func (m *DataModel) ModifyTargetEntity(target TargetEntity) error {
	if _, ok := m.targets[target.ID()]; ok {
		m.targets[target.ID()] = target
		return nil
	}
	return fmt.Errorf("target entity %s not found", target.ID())
}

// DeleteTargetEntity deletes an existing target entity.
func (m *DataModel) DeleteTargetEntity(target TargetEntity) error {
	if _, ok := m.targets[target.ID()]; ok {
		delete(m.targets, target.ID())
		return nil
	}
	return fmt.Errorf("target entity %s does not exist", target.ID())
}
