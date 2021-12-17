package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrdering_Format(t *testing.T) {
	type OrderingCase struct {
		Name    string
		AltName string
		Ordering
	}

	var cases = []OrderingCase{
		{"name", "name", Ordering{Vals: []*Ordering_Value{{Field: "name"}}}},
		{"name asc", "name   asc", Ordering{Vals: []*Ordering_Value{{Field: "name", Sort: Ordering_SORT_ASC}}}},
		{"name asc, count desc", "name asc   , count desc", Ordering{Vals: []*Ordering_Value{{Field: "name", Sort: Ordering_SORT_ASC}, {Field: "count", Sort: Ordering_SORT_DESC}}}},
		{"name, count desc", "name    , count desc", Ordering{Vals: []*Ordering_Value{{Field: "name"}, {Field: "count", Sort: Ordering_SORT_DESC}}}},
	}

	for _, c := range cases {
		o, err := ParseOrdering(c.AltName)
		assert.NoError(t, err)

		assert.Equal(t, len(c.Vals), len(o.Vals))
		for i := range o.Vals {
			assert.Equal(t, c.Vals[i].Field, o.Vals[i].Field)
			assert.Equal(t, c.Vals[i].Sort, o.Vals[i].Sort)
		}

		assert.Equal(t, c.Name, c.Format())
	}
}
