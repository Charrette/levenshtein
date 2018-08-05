package levenshtein

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	testCases := []struct {
		index            int
		s1               string
		s2               string
		opts             []Option
		expectedDistance int
	}{
		{0, "", "", nil, 0},
		{1, "foobar", "foobar", nil, 0},
		{2, "pétèràôlï", "peteraoli", nil, 10},
		{3, "pétèràôlï", "peteraoli", []Option{IgnoreDiacritics}, 0},
		{4, "noix de pécan", "noix de pecan", nil, 2},
		{5, "noix de pécan", "noix de pecan", []Option{IgnoreDiacritics}, 0},
		{6, "NoIx dE pécaN", "nOiX de pEcAn", nil, 9},
		{7, "NoIx dE pécaN", "nOiX de pEcAn", []Option{IgnoreCase}, 2},
		{8, "NoIx dE pécaN", "nOiX de pEcAn", []Option{IgnoreDiacritics}, 8},
		{9, "NoIx dE pécaN", "nOiX de pEcAn", []Option{IgnoreCase, IgnoreDiacritics}, 0},
		{10, "hello", "world", nil, 4},
		{11, "foo", "bar", nil, 3},
		{12, "baguette", "champignon", nil, 9},
		{13, "smartphone", "computer", nil, 9},
		{14, "tomate", "tomote", nil, 1},
		{15, "cucumber", "cucumb", nil, 2},
		{16, "cheese", "chaase", nil, 2},
		{17, "i'm the best", "i'm the worst", nil, 3},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("compare [%s] and [%s] %d", tc.s1, tc.s2, tc.index)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedDistance, Distance(tc.s1, tc.s2, tc.opts...))
		})
	}
}
