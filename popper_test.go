package popper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPopper(t *testing.T) {
	t.Run("Pop First", func(t *testing.T) {
		t.Run("Should Pass", func(t *testing.T) {
			t.Run("struct popper", func(t *testing.T) {
				p := New([]struct{ field string }{
					{"first"},
					{"last"},
				})
				got, err := p.PopFirst()
				require.NoError(t, err)
				require.Equal(t, struct{ field string }{"first"}, got)
			})
		})
		t.Run("Should Fail", func(t *testing.T) {
			t.Run("Empty Elements", func(t *testing.T) {
				p := New([]bool{true})
				e, err := p.PopFirst()
				require.NoError(t, err)
				require.True(t, e)
				// try popping the first element again
				// now that the underlying element collection is empty.
				_, err = p.PopFirst()
				require.Error(t, err)
				require.Equal(t, ErrEmptyElements, err)
			})
		})
	})
	t.Run("Pop Last", func(t *testing.T) {
		t.Run("Should Pass", func(t *testing.T) {
			t.Run("struct popper", func(t *testing.T) {
				p := New([]struct{ field string }{
					{"first"},
					{"last"},
				})
				got, err := p.PopLast()
				require.NoError(t, err)
				require.Equal(t, struct{ field string }{"last"}, got)
			})
		})
		t.Run("Should Fail", func(t *testing.T) {
			t.Run("Empty Elements", func(t *testing.T) {
				p := New([]bool{true})
				e, err := p.PopLast()
				require.NoError(t, err)
				require.True(t, e)
				// try popping the last element again
				// now that the underlying element collection is empty.
				_, err = p.PopLast()
				require.Error(t, err)
				require.Equal(t, ErrEmptyElements, err)
			})
		})
	})
	t.Run("Pop Element", func(t *testing.T) {
		t.Run("Should Pass", func(t *testing.T) {
			t.Run("struct popper", func(t *testing.T) {
				p := New([]struct{ field string }{
					{"first"},
					{"second"},
					{"last"},
				})
				require.NoError(t, p.PopElement(struct{ field string }{"second"}))
				require.Len(t, p.Elements(), 2)
				require.Equal(t, struct{ field string }{"last"}, p.Elements()[1])
			})
		})
		t.Run("Should Fail", func(t *testing.T) {
			t.Run("Empty Elements", func(t *testing.T) {
				p := New([]struct{ field string }{{"onlyone"}})
				first, err := p.PopFirst()
				require.NoError(t, err)
				require.Equal(t, struct{ field string }{"onlyone"}, first)
				err = p.PopElement(struct{ field string }{"doesntmatter"})
				require.Error(t, err)
				require.Equal(t, ErrEmptyElements, err)
			})
			t.Run("Element Not Found", func(t *testing.T) {
				p := New([]struct{ field string }{
					{"first"},
					{"second"},
					{"last"},
				})
				err := p.PopElement(struct{ field string }{"doesntexist"})
				require.Error(t, err)
				require.Equal(t, ErrElementNotFound, err)
			})
		})
	})
	t.Run("Pop Index", func(t *testing.T) {
		t.Run("Should Pass", func(t *testing.T) {
			p := New([]struct{ field string }{
				{"first"},
				{"second"},
				{"last"},
			})
			middleElement, err := p.PopIndex(1)
			require.NoError(t, err)
			require.Equal(t, struct{ field string }{"second"}, middleElement)
			require.Len(t, p.Elements(), 2)
		})
		t.Run("Should Fail", func(t *testing.T) {
			t.Run("Empty Elements", func(t *testing.T) {
				p := New([]struct{ field string }{{"onlyone"}})
				first, err := p.PopFirst()
				require.NoError(t, err)
				require.Equal(t, struct{ field string }{"onlyone"}, first)
				_, err = p.PopIndex(0)
				require.Error(t, err)
				require.Equal(t, ErrEmptyElements, err)
			})
			t.Run("Index Out Of Bounds", func(t *testing.T) {
				p := New([]struct{ field string }{
					{"first"},
					{"second"},
					{"last"},
				})
				_, err := p.PopIndex(3)
				require.Error(t, err)
				require.Equal(t, ErrIndexOutOfBounds, err)
			})
		})
	})
}
