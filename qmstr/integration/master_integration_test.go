package integration

import (
	"os"
	model "qmstr-prototype/qmstr/qmstr-model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandleSourceEntities(t *testing.T) {
	client := model.NewClient(MasterAddress)
	source := model.SourceEntity{Path: "src/hello.c", Hash: ""}
	// retrieve source entity:
	value, err := client.GetSourceEntity(source.ID())
	require.NoError(t, err, "error in GetSourceEntity")
	if len(value.ID()) != 0 {
		err = client.DeleteSourceEntity(value)
		require.NoError(t, err, "Deleting an existing entity should not cause an error.")
	}

	err = client.AddSourceEntity(source)
	require.NoError(t, err, "Creating a source entity that does not exist should not cause an error.")
	_, err = client.GetSourceEntity(source.ID())
	require.NoError(t, err, "Error in GetSourceEntity after adding the entity")
	source.Hash = "1234567890"
	err = client.ModifySourceEntity(source)
	require.NoError(t, err, "Error in ModifySourceEntity after adding the entity")
	value2, err := client.GetSourceEntity(source.ID())
	require.NoError(t, err, "Error in GetSourceEntity after modifying the entity")
	require.Equal(t, source.Hash, value2.Hash, "Returned value does not match what we set")
	err = client.DeleteSourceEntity(source)
	require.NoError(t, err, "Error deleting existing source entity")
	value3, err := client.GetSourceEntity(source.ID())
	require.Empty(t, value3.ID(), "The source entity I just deleted should not exist")
}

func TestHandleDependencyEntities(t *testing.T) {
	client := model.NewClient(MasterAddress)
	dep := model.DependencyEntity{Name: "libc.so", Hash: ""}
	if value, err := client.GetDependencyEntity(dep.ID()); len(value.ID()) != 0 {
		require.NoError(t, err, "error in DependencyEntity")
		err = client.DeleteDependencyEntity(value)
		require.NoError(t, err, "Deleting an existing entity should not cause an error.")
	}
	require.NoError(t, client.AddDependencyEntity(dep),
		"Creating a dependency entity that does not exist should not cause an error.")
	if value, err := client.GetDependencyEntity(dep.ID()); true {
		require.NoError(t, err, "Error in GetDependencyEntity after adding the entity")
		require.Equal(t, dep.ID(), value.ID(), "Saved and retrieved value should be equal")
	}
	dep.Hash = "1234567890"
	require.NoError(t, client.ModifyDependencyEntity(dep), "Error in ModifyDependencyEntity after adding the entity")
	if value, err := client.GetDependencyEntity(dep.ID()); true {
		require.NoError(t, err, "Error in GetDependencyEntity after modifying the entity")
		require.Equal(t, dep.Hash, value.Hash, "Returned value does not match what we set")
	}
	require.NoError(t, client.DeleteDependencyEntity(dep), "Error deleting existing dependency entity")
	if value, err := client.GetDependencyEntity(dep.ID()); true {
		require.NoError(t, err, "Querying a non-existant entity is not an error")
		require.Empty(t, value.ID(), "The dependency entity I just deleted should not exist")
	}
}

func TestHandleTargetEntities(t *testing.T) {
	client := model.NewClient(MasterAddress)
	tgt := model.TargetEntity{Name: "libc.so", Hash: ""}
	value, err := client.GetTargetEntity(tgt.ID())
	require.NoError(t, err, "GetTargetEntity should not fail")
	if len(value.ID()) != 0 {
		require.NoError(t, err, "error in TargetEntity")
		err = client.DeleteTargetEntity(value)
		require.NoError(t, err, "Deleting an existing entity should not cause an error.")
	}
	require.NoError(t, client.AddTargetEntity(tgt),
		"Creating a target entity that does not exist should not cause an error.")
	if value, err := client.GetTargetEntity(tgt.ID()); true {
		require.NoError(t, err, "Error in GetTargetEntity after adding the entity")
		require.Equal(t, tgt.ID(), value.ID(), "Saved and retrieved value should be equal")
	}
	tgt.Hash = "1234567890"
	require.NoError(t, client.ModifyTargetEntity(tgt), "Error in ModifyTargetEntity after adding the entity")
	if value, err := client.GetTargetEntity(tgt.ID()); true {
		require.NoError(t, err, "Error in GetTargetEntity after modifying the entity")
		require.Equal(t, tgt.Hash, value.Hash, "Returned value does not match what we set")
	}
	require.NoError(t, client.DeleteTargetEntity(tgt), "Error deleting existing target entity")
	if value, err := client.GetTargetEntity(tgt.ID()); true {
		require.NoError(t, err, "Querying a non-existant entity is not an error")
		require.Empty(t, value.ID(), "The target entity I just deleted should not exist")
	}
}

var MasterAddress string

func TestMain(m *testing.M) {
	// masterAddress := os.Getenv("QMSTR_MASTER_ADDRESS")
	masterAddress := "http://localhost:9000"
	if _, _, _, err := model.ParseMasterAddress(masterAddress); err != nil {
		panic(err)
	}
	MasterAddress = masterAddress
	os.Exit(m.Run())
}
