[![Go Reference](https://pkg.go.dev/badge/github.com/fuskovic/popper.svg)](https://pkg.go.dev/github.com/fuskovic/popper)
[![Go Report Card](https://goreportcard.com/badge/github.com/fuskovic/popper)](https://goreportcard.com/report/github.com/fuskovic/popper)
![CI](https://github.com/fuskovic/popper/actions/workflows/ci.yaml/badge.svg)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-97%25-brightgreen.svg?longCache=true&style=flat)</a>

# Popper

A minimalistic interface that provides various popping utilities on a generic slice of elements.

## Pre-requisites

Requires Go 1.18+

## Installation

    go get github.com/fuskovic/popper

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/fuskovic/popper"
)

type myCustomType struct {
	someString string
	someInt    int
}

func main() {
	elements := []myCustomType{
		{"first", 543},
		{"second", 234},
		{"third", 638},
		{"fourth", 143},
		{"fifth", 524},
		{"sixth", 611},
	}

	// Create a new popper.
	p := popper.New(elements)

	// PopFirst removes the first element from the underlying slice and returns it.
	firstElement, err := p.PopFirst()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("popped first element: %+v\n", firstElement)

	// PopLast removes the last element from the underlying slice and returns it.
	lastElement, err := p.PopLast()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("popped last element: %+v\n", lastElement)

	// PopElement removes the target element from the underlying slice.
	target := myCustomType{"third", 638}
	if err := p.PopElement(target); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("popped target element: %+v\n", target)
	fmt.Printf("number of elements: %d\n", p.Len())

	// PopIndex removes the element at the specified index from the underlying slice and returns it.
	element, err := p.PopIndex(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("popped element at index 2: %+v\n", element)

	// Get the number of remaining elements.
	fmt.Printf("number of elements: %d\n", p.Len())

	// You can use the Elements method to ditch the Popper interface when you're done 
	// with it and reassign the modified slice of elements to the original type.
	elements = p.Elements()
	fmt.Printf("checking elements: %v\n", elements)
}
```

## Output

```
popped first element: {someString:first someInt:543}
popped last element: {someString:sixth someInt:611}
popped target element: {someString:third someInt:638}
number of elements: 3
popped element at index 2: {someString:fifth someInt:524}
number of elements: 2
checking elements: [{second 234} {fourth 143}]
```