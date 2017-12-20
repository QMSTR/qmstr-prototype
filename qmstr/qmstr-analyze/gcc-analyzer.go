package analyze

import (
	"path/filepath"
	model "qmstr-prototype/qmstr/qmstr-model"
	"regexp"
	"strings"
)

type mode int

const (
	LINK mode = iota
	PREPROC
	COMPILE
	ASSEMBLE
)

var (
	sourceCodeExtensions = map[string]struct{}{
		".c":   struct{}{},
		".cpp": struct{}{},
		".c++": struct{}{},
		".cc":  struct{}{},
	}

	// regular expressions to extract libraries and library paths from the commandline
	linkLibPattern     = regexp.MustCompile(`^-l\s*(\S+)\s*`)
	linkLibPathPattern = regexp.MustCompile(`^-L\s*(\S+)\s*`)
)

// GNUCAnalyzer holds the analysis data
type GNUCAnalyzer struct {
	args             []string
	sources          []int
	target           []string
	mode             mode
	libs             []string
	libPath          []string
	givenTargetIndex int
}

// NewGNUCAnalyzer returns an initialized Analyzer to analyze gcc
func NewGNUCAnalyzer(args []string, debug bool) *GNUCAnalyzer {
	initLogging(debug)
	a := GNUCAnalyzer{args, []int{}, []string{}, LINK, []string{}, []string{"/usr/lib", "/usr/lib32", "/usr/lib64"}, 0}
	return &a
}

// Print will print the results of the command line analysis if in running in debug mode
func (a *GNUCAnalyzer) Print() {
	Logger.Printf("The source files are:")
	for _, arg := range a.sources {
		Logger.Printf(a.args[arg])
	}
	Logger.Printf("The targets are:")
	for _, arg := range a.target {
		Logger.Printf(arg)
	}
	if a.mode == LINK {
		Logger.Printf("The libraries are:")
		for _, arg := range a.libs {
			Logger.Printf(arg)
		}
	}
}

// Analyze will analyze the command line parameters and detect source files, targets and linked libraries.
func (a *GNUCAnalyzer) Analyze(simulate bool) *GNUCAnalyzer {
	a.extractMode()
	a.detectSourceFiles()
	a.detectTarget()
	if a.mode == LINK {
		a.extractLibs()
		a.detectObjectFiles()
	}
	return a
}

func (a *GNUCAnalyzer) extractMode() {
	for _, arg := range a.args {
		switch arg {
		case "-E":
			a.mode = PREPROC
		case "-S":
			a.mode = COMPILE
		case "-c":
			a.mode = ASSEMBLE
		}
	}
}

func (a *GNUCAnalyzer) detectTarget() {
	if len(a.target) != 0 {
		return
	}
	for index, arg := range a.args {
		if arg == "-o" {
			a.target = append(a.target, a.args[index+1])
			a.givenTargetIndex = len(a.target) - 1
		}
	}
	if len(a.target) == 0 {
		switch a.mode {
		case PREPROC:
			return
		case COMPILE:
			return
		case ASSEMBLE:
			for _, srcIdx := range a.sources {
				filename := a.args[srcIdx]
				objectname := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".o"
				a.target = append(a.target, objectname)
			}
		case LINK:
			a.target = append(a.target, "a.out")
		}
	}
}

func (a *GNUCAnalyzer) detectSourceFiles() {
	for index, arg := range a.args {
		if extension := filepath.Ext(arg); extension != "" {
			_, ok := sourceCodeExtensions[extension]
			if ok {
				a.sources = append(a.sources, index)
			}
		}
	}
}

func (a *GNUCAnalyzer) extractLibs() {
	for _, arg := range a.args {
		matches := linkLibPattern.FindStringSubmatch(arg)
		if matches != nil {
			a.libs = append(a.libs, matches[1])
		}

		matches = linkLibPathPattern.FindStringSubmatch(arg)
		if matches != nil {
			a.libPath = append(a.libPath, matches[1])
		}
	}
}

func (a *GNUCAnalyzer) detectObjectFiles() {
	for index, arg := range a.args {
		if extension := filepath.Ext(arg); extension == ".o" {
			if a.args[index-1] != "-o" {
				a.sources = append(a.sources, index)
			}
		}
	}
}

// SendResults will transmit the results of the analysis to the master server
func (a *GNUCAnalyzer) SendResults() {
	client := model.NewClient("http://localhost:9000/")
	if a.mode == ASSEMBLE {
		for idx, target := range a.target {
			var t model.TargetEntity
			t.Name = target
			t.Hash = "targethash"
			t.Linked = false

			var s model.SourceEntity
			s.Path = a.args[a.sources[idx]]
			s.Hash = "filehash"
			s.Licenses = analyzeSourceFile(s.Path)
			client.AddSourceEntity(s)

			t.Sources = []string{s.ID()}
			client.AddTargetEntity(t)
		}
	} else if a.mode == LINK {
		var t model.TargetEntity
		t.Name = a.target[a.givenTargetIndex]
		t.Hash = "linktargethash"
		t.Linked = true

		for _, srcIdx := range a.sources {
			var s model.SourceEntity
			s.Path = a.args[srcIdx]
			s.Hash = "filehash"
			client.AddSourceEntity(s)
			t.Sources = append(t.Sources, s.ID())
		}

		for _, lib := range a.libs {
			var d model.DependencyEntity
			d.Name = lib
			d.Hash = "dephash"
			client.AddDependencyEntity(d)
			t.Dependencies = append(t.Dependencies, d.ID())
		}
		client.AddTargetEntity(t)
	}
}
