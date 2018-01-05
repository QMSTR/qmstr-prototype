package model

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCRUDSourceEntity(t *testing.T) {
	source := SourceEntity{"a/b/source.c", "1234567890", []string{"GPL-3.0"}}
	//make sure the value is not in the model:
	if value, err := dataModel.GetSourceEntity(source.ID()); err != nil {
		dataModel.DeleteSourceEntity(value)
	}
	_, err := dataModel.GetSourceEntity(source.ID())
	require.Error(t, err, "After removing it, the source entity should not exist")
	//Add the value, make sure it is there:
	dataModel.AddSourceEntity(source)
	value, err := dataModel.GetSourceEntity(source.ID())
	require.NoError(t, err, "After removing it, the source entity should not exist")
	require.Equal(t, source, value)
	value.Hash = "0987654321"
	err2 := dataModel.ModifySourceEntity(value)
	require.NoError(t, err2, "There should be no error updating an existing entity")
	value2, err3 := dataModel.GetSourceEntity(source.ID())
	require.NoError(t, err3, "After removing it, the source entity should not exist")
	require.Equal(t, value, value2)
	err4 := dataModel.DeleteSourceEntity(source)
	require.NoError(t, err4, "There should be no error deleting an existing source entity")
	value3, err5 := dataModel.GetSourceEntity(source.ID())
	require.Error(t, err5, "After removing it, the source entity should not exist")
	require.Equal(t, "", value3.ID())
}

func TestCRUDDependencyEntity(t *testing.T) {
	dependency := DependencyEntity{[]*TargetEntity{}, "/usr/lib/libc.so", "1234567890"}
	//make sure the value is not in the model:
	if value, err := dataModel.GetDependencyEntity(dependency.ID()); err != nil {
		dataModel.DeleteDependencyEntity(value)
	}
	_, err := dataModel.GetDependencyEntity(dependency.ID())
	require.Error(t, err, "After removing it, the dependency entity should not exist")
	//Add the value, make sure it is there:
	dataModel.AddDependencyEntity(dependency)
	value, err := dataModel.GetDependencyEntity(dependency.ID())
	require.NoError(t, err, "After removing it, the dependency entity should not exist")
	require.Equal(t, dependency, value)
	value.Hash = "0987654321"
	err2 := dataModel.ModifyDependencyEntity(value)
	require.NoError(t, err2, "There should be no error updating an existing entity")
	value2, err3 := dataModel.GetDependencyEntity(dependency.ID())
	require.NoError(t, err3, "After removing it, the dependency entity should not exist")
	require.Equal(t, value, value2)
	err4 := dataModel.DeleteDependencyEntity(dependency)
	require.NoError(t, err4, "There should be no error deleting an existing dependency entity")
	value3, err5 := dataModel.GetDependencyEntity(dependency.ID())
	require.Error(t, err5, "After removing it, the dependency entity should not exist")
	require.Equal(t, "", value3.ID())
}
func TestCRUDTargetEntity(t *testing.T) {
	target := TargetEntity{"bin/application", "1234567890", "", nil, nil, true}
	//make sure the value is not in the model:
	if value, err := dataModel.GetTargetEntity(target.ID()); err != nil {
		dataModel.DeleteTargetEntity(value)
	}
	_, err := dataModel.GetTargetEntity(target.ID())
	require.Error(t, err, "After removing it, the target entity should not exist")
	//Add the value, make sure it is there:
	dataModel.AddTargetEntity(target)
	value, err := dataModel.GetTargetEntity(target.ID())
	require.NoError(t, err, "After removing it, the target entity should not exist")
	require.Equal(t, target, value)
	value.Hash = "0987654321"
	err2 := dataModel.ModifyTargetEntity(value)
	require.NoError(t, err2, "There should be no error updating an existing entity")
	value2, err3 := dataModel.GetTargetEntity(target.ID())
	require.NoError(t, err3, "After removing it, the target entity should not exist")
	require.Equal(t, value, value2)
	err4 := dataModel.DeleteTargetEntity(target)
	require.NoError(t, err4, "There should be no error deleting an existing target entity")
	value3, err5 := dataModel.GetTargetEntity(target.ID())
	require.Error(t, err5, "After removing it, the target entity should not exist")
	require.Equal(t, "", value3.ID())
}

var dataModel *DataModel

func TestMain(m *testing.M) {
	dataModel = NewModel()
	os.Exit(m.Run())
}
