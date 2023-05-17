package oracle

import (
	"math"
	"math/rand"
	"testing"

	"github.com/kendfss/oprs"
	"github.com/kendfss/oracle/slices"
)

func TestSum(t *testing.T) {
	type test struct {
		terms    []int
		expected int
	}
	tests := []test{
		{terms: []int{0}, expected: 0},
		{terms: []int{1}, expected: 1},
		{terms: []int{0, 1}, expected: 1},
		{terms: []int{1, 2, 3}, expected: 6},
		{terms: []int{4, 5, 6, 7}, expected: 22},
	}
	for i, test := range tests {
		if s := slices.Reduce(oprs.Add[int], test.terms); s != test.expected {
			Quitf(t, "#%d (%v): should be %d but is %d", i, test.terms, test.expected, s)
		}
	}
}

func TestMkr(t *testing.T) {
	const (
		nTests = 10
		nItems = 20
	)

	for i := 0; i < nTests; i++ {
		result := Mkr(nItems, math.MaxInt)
		if len(result) != nItems {
			Quitf(t, "#%d (%v): has %d items, should have %d", i, result, len(result), nItems)
		}
		if shrink := slices.Compacted(result); len(result) != len(shrink) {
			Quitf(t, "#%d (%v): contains duplicates", i, result)
		}
	}
}

func TestMk(t *testing.T) {
	type argSet struct {
		len, quan int
	}
	const (
		testQuan = 10
		testLen  = 20
	)
	tests := [testQuan]argSet{}
	for i := 0; i < testQuan; i++ {
		tests[i] = argSet{len: rand.Intn(testLen), quan: rand.Intn(testQuan)}
	}
	for i, test := range tests {
		result := Mk(test.len, test.quan, 30)
		if len(result) != test.quan {
			Quitf(t, "#%d (%v): has %d items, should have %d", i, result, len(result), test.quan)
		}
		for j, wedge := range result {
			if len(wedge) != test.len {
				Quitf(t, "#%d.%d (%v.%v): has %d items, should have %d", i, j, result, wedge, len(wedge), test.len)
			}
		}
	}
}

func TestSubtractions(t *testing.T) {
	const nTests = 21
	indices := slices.Rangen(nTests)
	indices = append(indices, nTests, nTests+1, nTests+2)
	for i, e := range indices {
		result := Subtractions(nTests, e)
		if sum := slices.Reduce(oprs.Add[int], result); sum != nTests {
			Quitf(t, "test #%d (%d) yielded %d instead of %d", i, e, sum, nTests)
		}
	}
}
