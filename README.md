# Dynamic Slicer

The `dynamic-slicer` command is a code generator that will generate a function
for a type that will automatically manage resizing of a slice when attempting to
set a value at a specific index.

## Concept

If you have a `[]int` slice created with a command such as `make([]int, 0])`,
then you will run into `index out of range` errors if you attempt to insert
a value into index `1`.

For example

```Go
m := make([]int, 0])
m[1] := 5 // this is a runtime error.
```

The `dynamic-slicer` tool will generate a function that accepts the slice of the
specified type, the target index, and the value for that index, and will
automatically resize the slice (if needed), or simply set the value if not.

## Usage

Simply add a `go generate` magic comment to any source file.

```Go
//go:generate dynamic-slicer MyCustomType
```

Then run `go generate ./...` or equivalent.

## Example

To see this in action:

- Clone this repository (e.g. `git clone https://github.com/komish/dynamic-slicer`)
- Install the binary on your system and add it to your path (e.g. `go build .&& mv dynamic-slicer /usr/local/bin/`)
- Move into the `demo` directory: (e.g. `cd demo`)
- Run the generator (e.g. `go generate ./...`)
- Test the demo exection (which relies on this function existing) (e.g. `go run .` from within `demo/`)

## Footnotes

This project was an exercise in studying code generation in Golang.
