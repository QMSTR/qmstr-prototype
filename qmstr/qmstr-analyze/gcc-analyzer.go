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

type GNUCAnalyzer struct {
	args             []string
	sources          []int
	target           []string
	stash            []int
	mode             mode
	libs             []string
	libPath          []string
	givenTargetIndex int
}

// NewGNUCAnalyzer returns an initialized Analyzer to analyze gcc
func NewGNUCAnalyzer(args []string) *GNUCAnalyzer {
	a := GNUCAnalyzer{args, []int{}, []string{}, []int{}, LINK, []string{}, []string{"/usr/lib", "/usr/lib32", "/usr/lib64"}, -1}
	return &a
}

func (a *GNUCAnalyzer) Print() {
	Info.Printf("The source files are:")
	for _, arg := range a.sources {
		Info.Printf(a.args[arg])
	}
	Info.Printf("The targets are:")
	for _, arg := range a.target {
		Info.Printf(arg)
	}
	if a.mode == LINK {
		Info.Printf("The libraries are:")
		for _, arg := range a.libs {
			Info.Printf(arg)
		}
	}
}

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

func (a *GNUCAnalyzer) detectSourceFiles() *GNUCAnalyzer {

	for index, arg := range a.args {
		if extension := filepath.Ext(arg); extension != "" {
			_, ok := sourceCodeExtensions[extension]
			if ok {
				a.sources = append(a.sources, index)
			}
		}
	}
	return a
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

func (a *GNUCAnalyzer) detectObjectFiles() *GNUCAnalyzer {
	for index, arg := range a.args {
		if extension := filepath.Ext(arg); extension == ".o" {
			if a.args[index-1] != "-o" {
				a.sources = append(a.sources, index)
			}
		}
	}
	return a
}

func (a *GNUCAnalyzer) SendResults() {
	client := model.NewClient("http://localhost:8080/")
	if a.mode == ASSEMBLE {
		for idx, target := range a.target {
			var t model.TargetEntity
			t.Name = target
			t.Hash = "targethash"

			var s model.SourceEntity
			s.Path = a.args[a.sources[idx]]
			s.Hash = "filehash"
			s.Licenses = AnalyzeSourceFile(s.Path)
			client.AddSourceEntity(s)

			t.Sources = []string{s.ID()}
			client.AddTargetEntity(t)
		}
	} else if a.mode == LINK {
		var t model.TargetEntity
		t.Name = a.target[a.givenTargetIndex]
		t.Hash = "linktargethash"

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
