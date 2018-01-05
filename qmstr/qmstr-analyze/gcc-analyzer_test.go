package analyze

import (
	"testing"
)

func TestGCCAnalyzerArguments(t *testing.T) {
	args := []string{"-c", "array.c", "-o", "array.o"}
	// create a Canalyzer struct with the arguments
	analyzer := NewGNUCAnalyzer(args, false)
	analyzer.Analyze(true)

	//check if qmstr filled the struct with the correct values
	value := analyzer.args[analyzer.sources[0]]
	if value != "array.c" {
		t.Error("Expecting the value 'array.c' but got: ", value)
	}
	target := analyzer.target[0]
	if target != "array.o" {
		t.Error("Expecting the value 'array.o' but got: ", target)
	}
}

func TestGCCAnalyzerWithoutTarget(t *testing.T) {
	args := []string{"-c", "array.c"}
	// create a Canalyzer struct with the arguments
	analyzer := NewGNUCAnalyzer(args, false)
	analyzer.Analyze(true)

	//check if qmstr filled the struct with the correct values
	value := analyzer.args[analyzer.sources[0]]
	if value != "array.c" {
		t.Error("Expecting the value 'array.c' but got: ", value)
	}
	target := analyzer.target[0]
	if target != "array.o" {
		t.Error("Expecting the value 'array.o' but got: ", target)
	}
}

func TestGCCAnalyzerLinkWithoutTarget(t *testing.T) {
	args := []string{"array.c"}
	// create a Canalyzer struct with the arguments
	analyzer := NewGNUCAnalyzer(args, false)
	analyzer.Analyze(true)

	//check if qmstr filled the struct with the correct values
	value := analyzer.args[analyzer.sources[0]]
	if value != "array.c" {
		t.Error("Expecting the value 'array.c' but got: ", value)
	}
	target := analyzer.target[0]
	if target != "a.out" {
		t.Error("Expecting the value 'array.o' but got: ", target)
	}
}

func TestGCCAnalyzerMultiSourceArguments(t *testing.T) {
	args := []string{"array.c", "foo.c", "bar.c", "-o", "theprogram"}
	// create a Canalyzer struct with the arguments
	analyzer := NewGNUCAnalyzer(args, false)
	analyzer.Analyze(true)

	//check if qmstr filled the struct with the correct values
	value := analyzer.args[analyzer.sources[0]]
	if value != "array.c" {
		t.Error("Expecting the value 'array.c' but got: ", value)
	}
	value = analyzer.args[analyzer.sources[1]]
	if value != "foo.c" {
		t.Error("Expecting the value 'foo.c' but got: ", value)
	}
	value = analyzer.args[analyzer.sources[2]]
	if value != "bar.c" {
		t.Error("Expecting the value 'bar.c' but got: ", value)
	}
	target := analyzer.target[0]
	if target != "theprogram" {
		t.Error("Expecting the value 'theprogram' but got: ", target)
	}
}
