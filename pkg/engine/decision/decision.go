package decision

import "fmt"

// Instead of using a hashmap we implement our own stack based set.
type Decision struct {
	set  [5]bool
	size uint

	// A Player can decide to roll all dices if they all marked.
	All bool
}

func From(items ...uint) (*Decision, error) {
	decision := New()

	if err := decision.Insert(items...); err != nil {
		return nil, err
	}

	return decision, nil
}

func New() *Decision {
	return &Decision{}
}

func (d *Decision) Insert(items ...uint) error {
	for _, item := range items {
		if item > 4 {
			return fmt.Errorf("%d is strictly greater than 4", item)
		}

		if !d.DoesContain(item) {
			d.size += 1
			d.set[item] = true
		}
	}

	return nil
}

func (d Decision) IsEmpty() bool {
	return d.size == 0
}

func (d Decision) DoesContain(item uint) bool {
	return d.set[item]
}

func (d *Decision) toSlice() []uint {
	slice := make([]uint, 0, d.size)

	for index, ok := range d.set {
		if ok {
			slice = append(slice, uint(index))
		}
	}

	return slice
}

func (d Decision) Size() uint {
	return d.size
}

func (d Decision) String() string {
	return fmt.Sprint(d.toSlice())
}
