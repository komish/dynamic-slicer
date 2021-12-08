package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

const goFileEnv = "GOFILE"
const goPackageEnv = "GOPACKAGE"
const bugURL = "https://github.com/komish/dynamic-slicer/issues"

var help = fmt.Sprintf(`This generator expects exactly one argument with the
exact case used to describe a single type in your source code,
along with these environment variables commonly made available when calling
"go generate"

%s %s

E.g. 
    %s MyCustomType
	
File any bugs at: %s`, goFileEnv, goPackageEnv, os.Args[0], bugURL)

type SourceCode struct {
	Type           string
	File           string
	Package        string
	NormalizedType string
}

func (sc *SourceCode) isEmpty() bool {
	return !(len(sc.Type) != 0 &&
		len(sc.File) != 0 &&
		len(sc.Package) != 0)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(help)
		os.Exit(1)
	}

	normalizedType, err := validateAndNormalizeType(os.Args[1])
	if err != nil {
		fmt.Println("Provided type container unsupported characters other than a leading asterisk:", os.Args[1])
		os.Exit(2)

	}

	sourceCodeMeta := SourceCode{
		Type:           os.Args[1],
		File:           os.Getenv(goFileEnv),
		Package:        os.Getenv(goPackageEnv),
		NormalizedType: normalizedType,
	}

	if sourceCodeMeta.isEmpty() {
		fmt.Println("Some input value was empty")
		fmt.Println(help)
		os.Exit(3)
	}

	// Read in the template
	tpl := template.Must(template.New("generatedSource").Parse(funcTemplateText))

	// Prepare the type to include in the filename.
	typeLowerCase := strings.ToLower(sourceCodeMeta.NormalizedType)
	sourceFileWithoutExt := strings.Split(sourceCodeMeta.File, ".")[0]

	generatedFile := fmt.Sprintf("%s_%s_%s_gen.go", sourceFileWithoutExt, typeLowerCase, "setter")
	f, err := os.Create(generatedFile)
	if err != nil {
		// TODO: improve error handling
		fmt.Println("Could not create the output file:", generatedFile)
		panic(err)
	}

	b := bufio.NewWriter(f)
	err = tpl.Execute(b, sourceCodeMeta)
	if err != nil {
		// TODO: improve error handling
		panic(err)
	}

	err = b.Flush()
	if err != nil {
		// TODO: improve error handling
		panic(err)
	}
}

// validateAndNormalizeType accepts the type string provided and will convert a leading
// asterisk (*) into alpha characters for use in source code. If any other special characters
// are found, this will throw an error.
func validateAndNormalizeType(t string) (string, error) {
	normalizedT := t
	first := t[0:1] // check for pointer
	if first == "*" {
		normalizedT = t[1:] + "Ptr"
	}

	if !regexp.MustCompile(`^[A-Za-z]+$`).MatchString(normalizedT) {
		return "", errors.New("unsupported characters in type")
	}

	return normalizedT, nil
}
