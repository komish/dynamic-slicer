package main

var funcTemplateText string = `package {{ .Package }}

// This file is generated. Do not edit manually.
// Input file: {{ .File }}.

// Set{{ .NormalizedType }}ValueAtIndex will set the provided value at the provided index in the
// provided slice. If the slice is not the appropriate length to support the index, the slice
// will be extended to accomodate the index.
func Set{{ .NormalizedType }}ValueAtIndex(slice []{{ .Type }}, value {{.Type}}, index int) []{{ .Type }} {
	sliceLengthAtIndex := index + 1

	if len(slice) == 0 {
		slice = make([]{{ .Type }}, sliceLengthAtIndex)
		slice[index] = value
		return slice
	}

	sliceLength := len(slice)
	if sliceLength >= sliceLengthAtIndex {
		slice[index] = value
		return slice
	}

	sliceLengthDiff := sliceLengthAtIndex - sliceLength
	extension := make([]{{ .Type }}, sliceLengthDiff)
	extension[sliceLengthDiff-1] = value
	return append(slice, extension...)
}`
