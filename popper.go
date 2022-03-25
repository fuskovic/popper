package popper

import "errors"

var (
	// ErrEmptyElements is returned when a pop operation is attempted on an empty set of elements.
	ErrEmptyElements = errors.New("empty elements")
	// ErrElementNotFound is returned when the element designated to pop out of the collection is not found.
	ErrElementNotFound = errors.New("element not found")
	// ErrIndexOutOfBounds is returned when the index specified exceeds the length of the underlying collection of elements.
	ErrIndexOutOfBounds = errors.New("index out of bounds")
)

type (
	// Popper is an interface used for modifying a collection of elements.
	Popper[T comparable] interface {
		// PopFirst removes the first element from the underlying collection and returns the removed element.
		PopFirst() (T, error)

		// PopLast removes the last element from the underlying collection and returns the removed element.
		PopLast() (T, error)

		// PopElement removes element from the collection. If element does not exist in the collection, a popper.ErrElementNotFound is returned.
		PopElement(element T) error

		// PopIndex removes the element at index in the collection and returns the removed element.
		// If the index exceeds the length of the collection of elements, a popper.ErrIndexOutOfBounds is returned.
		PopIndex(index int) (T, error)

		// Elements returns the underlying collection of elements.
		Elements() []T

		// Len returns the number of elements.
		Len() int
	}
	popper[T comparable] struct{ elements []T }
)

// New initializes a new Popper for a generic collection of elements.
func New[T comparable](elements []T) (Popper[T], error) {
	if len(elements) == 0 {
		return nil, ErrEmptyElements
	}
	i := new(popper[T])
	i.elements = elements
	return i, nil
}

func (p *popper[T]) PopFirst() (T, error) {
	var first T
	if len(p.elements) == 0 {
		return first, ErrEmptyElements
	}
	first = p.elements[0]
	p.elements = p.elements[1:]
	return first, nil
}

func (p *popper[T]) PopLast() (T, error) {
	var last T
	if len(p.elements) == 0 {
		return last, ErrEmptyElements
	}
	last = p.elements[len(p.elements)-1]
	p.elements = p.elements[0 : len(p.elements)-1]
	return last, nil
}

func (p *popper[T]) PopElement(element T) error {
	if len(p.elements) == 0 {
		return ErrEmptyElements
	}

	var found bool
	for i := range p.elements {
		if p.elements[i] == element {
			found = true
			p.elements = append(p.elements[:i], p.elements[i+1:]...)
			break
		}
	}

	if !found {
		return ErrElementNotFound
	}
	return nil
}

func (p *popper[T]) PopIndex(index int) (T, error) {
	var element T
	if len(p.elements) == 0 {
		return element, ErrEmptyElements
	}
	if index > len(p.elements)-1 {
		return element, ErrIndexOutOfBounds
	}
	element = p.elements[index]
	p.elements = append(p.elements[:index], p.elements[index+1:]...)
	return element, nil
}

func (p *popper[T]) Elements() []T { return p.elements }
func (p *popper[T]) Len() int      { return len(p.elements) }
