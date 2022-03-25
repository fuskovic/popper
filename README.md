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
	firstElement, err := popper.PopFirst()
	if err != nil {
		log.Fatal(err)
	}

	// PopLast removes the last element from the underlying slice and returns it.
	lastElement, err := popper.PopLast()
	if err != nil {
		log.Fatal(err)
	}

	// PopElement removes the target element from the underlying slice.
	if err := popper.PopElement(customType{"third", 638}); err != nil {
		log.Fatal(err)
	}

	// PopIndex removes the element at the specified index from the underlying slice and returns it.
	element, err := popper.PopIndex(2)
	if err != nil {
		log.Fatal(err)
	}

	// List remaining elements.
	fmt.Printf("checking elements: %v\n", popper.Elements())

	// Get the number of remaining elements.
	fmt.Printf("number of elements: %d\n", popper.Len())
}
```