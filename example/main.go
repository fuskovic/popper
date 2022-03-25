package main

import (
	"fmt"
	"log"

	"github.com/fuskovic/popper"
)

type customType struct {
	someString string
	someInt    int
}

func main() {
	popper, err := popper.New(
		[]customType{
			{"doesntmatter", 543},
			{"doesntmatteragain", 234},
			{"stilldoesntmatter", 638},
			{"ohthenegligence", 143},
			{"lastone", 524},
			{"forrealthistime", 611},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("elements: %v\nlen: %d\n", popper.Elements(), popper.Len())

	firstElement, err := popper.PopFirst()
	if err != nil {
		log.Fatalf("failed to pop first: %s\n", err)
	}

	fmt.Printf(
		"popped first element: %+v\nremaining_elements: %v\nnum_remaining: %d\n",
		firstElement, popper.Elements(), popper.Len(),
	)

	lastElement, err := popper.PopLast()
	if err != nil {
		log.Fatalf("failed to pop last: %s\n", err)
	}

	fmt.Printf(
		"popped last element: %+v\nremaining_elements: %v\nnum_remaining: %d\n",
		lastElement, popper.Elements(), popper.Len(),
	)

	targetElement := customType{"stilldoesntmatter", 638}
	if err := popper.PopElement(targetElement); err != nil {
		log.Fatalf("failed to pop target element: %s\n", err)
	}

	fmt.Printf(
		"popped target element: %+v\nremaining_elements: %v\nnum_remaining: %d\n",
		targetElement, popper.Elements(), popper.Len(),
	)

	element, err := popper.PopIndex(2)
	if err != nil {
		log.Fatalf("failed to pop item at index 2: %v\n", err)
	}

	fmt.Printf(
		"popped element at index 2: %+v\nremaining_elements: %v\nnum_remaining: %d\n",
		element, popper.Elements(), popper.Len(),
	)
}
