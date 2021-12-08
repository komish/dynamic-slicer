package main

//go:generate dynamic-slicer MyCustomType
type MyCustomType struct {
	S string
}

//go:generate dynamic-slicer *SomeOtherType
type SomeOtherType struct {
	S string
}
