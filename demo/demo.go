package main

import "fmt"

var values []MyCustomType = []MyCustomType{
	{
		S: "Foo",
	},
	{
		S: "Bar",
	},
	{
		S: "Baz",
	},
	{
		S: "Noz",
	},
	{
		S: "Tos",
	},
}

var idxs []int = []int{1, 0, 3, 4, 2}

func main() {
	fmt.Println("Start")

	var m []MyCustomType

	for _, i := range idxs {
		m = SetMyCustomTypeValueAtIndex(m, values[i], i)
		fmt.Println(m)
		fmt.Println("---")
	}
}
