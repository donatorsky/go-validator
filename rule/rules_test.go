package rule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Bailer(t *testing.T) {
	// given
	bailer := &Bailer{}

	// then
	require.False(t, bailer.bailed)
	require.False(t, bailer.Bails())

	// and when
	bailer.MarkBailed()

	// then
	require.True(t, bailer.bailed)
	require.True(t, bailer.Bails())
	require.False(t, bailer.bailed)
}

func Test_Dereference(t *testing.T) {
	// given
	for ttIdx, tt := range []struct {
		in    any
		out   any
		isNil bool
	}{
		// nil
		{
			in:    nil,
			out:   nil,
			isNil: true,
		},

		// scalar value
		{
			in:    (*string)(nil),
			out:   nil,
			isNil: true,
		},
		{
			in:    ptr((*string)(nil)),
			out:   nil,
			isNil: true,
		},
		{
			in:    "foo",
			out:   "foo",
			isNil: false,
		},
		{
			in:    ptr("foo"),
			out:   "foo",
			isNil: false,
		},
		{
			in:    ptr(ptr("foo")),
			out:   "foo",
			isNil: false,
		},

		// slice
		{
			in:    ([]int)(nil),
			out:   nil,
			isNil: true,
		},
		{
			in:    ptr(([]int)(nil)),
			out:   nil,
			isNil: true,
		},
		{
			in:    []int{},
			out:   []int{},
			isNil: false,
		},
		{
			in:    ptr([]int{}),
			out:   []int{},
			isNil: false,
		},
		{
			in:    ptr(ptr([]int{})),
			out:   []int{},
			isNil: false,
		},

		// array
		{
			in:    [1]int{},
			out:   [1]int{},
			isNil: false,
		},
		{
			in:    ptr([1]int{}),
			out:   [1]int{},
			isNil: false,
		},
		{
			in:    ptr(ptr([1]int{})),
			out:   [1]int{},
			isNil: false,
		},

		// map
		{
			in:    (map[any]any)(nil),
			out:   nil,
			isNil: true,
		},
		{
			in:    ptr((map[any]any)(nil)),
			out:   nil,
			isNil: true,
		},
		{
			in:    map[any]any{},
			out:   map[any]any{},
			isNil: false,
		},
		{
			in:    ptr(map[any]any{}),
			out:   map[any]any{},
			isNil: false,
		},
		{
			in:    ptr(ptr(map[any]any{})),
			out:   map[any]any{},
			isNil: false,
		},
	} {
		t.Run(fmt.Sprintf("#%d", ttIdx), func(t *testing.T) {
			// when
			value, isNil := Dereference(tt.in)

			// then
			require.Equal(t, tt.out, value)
			require.Equal(t, tt.isNil, isNil)
		})
	}
}

func Test_CompareNumbers(t *testing.T) {
	// given
	for _, tt := range []compareNumbersTestCase[int, int]{
		{
			left:     5,
			right:    5,
			expected: 0,
		},
		{
			left:     4,
			right:    5,
			expected: -1,
		},
		{
			left:     5,
			right:    4,
			expected: 1,
		},
	} {
		runCompareNumbersTestCase(t, tt)
	}

	// and given
	for _, tt := range []compareNumbersTestCase[int, uint]{
		{
			left:     5,
			right:    5,
			expected: 0,
		},
		{
			left:     4,
			right:    5,
			expected: -1,
		},
		{
			left:     5,
			right:    4,
			expected: 1,
		},
	} {
		runCompareNumbersTestCase(t, tt)
	}

	// and given
	for _, tt := range []compareNumbersTestCase[int, float64]{
		{
			left:     5,
			right:    5,
			expected: 0,
		},
		{
			left:     4,
			right:    5,
			expected: -1,
		},
		{
			left:     5,
			right:    4,
			expected: 1,
		},
	} {
		runCompareNumbersTestCase(t, tt)
	}

	// and given
	for _, tt := range []compareNumbersTestCase[uint, int]{
		{
			left:     5,
			right:    5,
			expected: 0,
		},
		{
			left:     4,
			right:    5,
			expected: -1,
		},
		{
			left:     5,
			right:    4,
			expected: 1,
		},
	} {
		runCompareNumbersTestCase(t, tt)
	}

	// and given
	for _, tt := range []compareNumbersTestCase[uint, uint]{
		{
			left:     5,
			right:    5,
			expected: 0,
		},
		{
			left:     4,
			right:    5,
			expected: -1,
		},
		{
			left:     5,
			right:    4,
			expected: 1,
		},
	} {
		runCompareNumbersTestCase(t, tt)
	}

	// and given
	for _, tt := range []compareNumbersTestCase[uint, float64]{
		{
			left:     5,
			right:    5,
			expected: 0,
		},
		{
			left:     4,
			right:    5,
			expected: -1,
		},
		{
			left:     5,
			right:    4,
			expected: 1,
		},
	} {
		runCompareNumbersTestCase(t, tt)
	}

	// and given
	for _, tt := range []compareNumbersTestCase[float64, int]{
		{
			left:     5,
			right:    5,
			expected: 0,
		},
		{
			left:     4,
			right:    5,
			expected: -1,
		},
		{
			left:     5,
			right:    4,
			expected: 1,
		},
	} {
		runCompareNumbersTestCase(t, tt)
	}

	// and given
	for _, tt := range []compareNumbersTestCase[float64, uint]{
		{
			left:     5,
			right:    5,
			expected: 0,
		},
		{
			left:     4,
			right:    5,
			expected: -1,
		},
		{
			left:     5,
			right:    4,
			expected: 1,
		},
	} {
		runCompareNumbersTestCase(t, tt)
	}

	// and given
	for _, tt := range []compareNumbersTestCase[float64, float64]{
		{
			left:     5,
			right:    5,
			expected: 0,
		},
		{
			left:     4,
			right:    5,
			expected: -1,
		},
		{
			left:     5,
			right:    4,
			expected: 1,
		},
	} {
		runCompareNumbersTestCase(t, tt)
	}
}

func BenchmarkCompareNumbers(b *testing.B) {
	runCompareNumbersBenchmark[int, int](b, 5, 5)
	runCompareNumbersBenchmark[int, uint](b, 5, 5)
	runCompareNumbersBenchmark[int, float64](b, 5, 5)
	runCompareNumbersBenchmark[uint, int](b, 5, 5)
	runCompareNumbersBenchmark[uint, uint](b, 5, 5)
	runCompareNumbersBenchmark[uint, float64](b, 5, 5)
	runCompareNumbersBenchmark[float64, int](b, 5, 5)
	runCompareNumbersBenchmark[float64, uint](b, 5, 5)
	runCompareNumbersBenchmark[float64, float64](b, 5, 5)
}

func runCompareNumbersTestCase[N1, N2 numberType](t *testing.T, tt compareNumbersTestCase[N1, N2]) {
	t.Run(compareNumbersTestCaseName(tt.left, tt.right), func(t *testing.T) {
		// when
		result := CompareNumbers(tt.left, tt.right)

		// then
		require.Equal(t, tt.expected, result)
	})
}

func runCompareNumbersBenchmark[N1, N2 numberType](b *testing.B, left N1, right N2) {
	b.Run(compareNumbersTestCaseName[N1, N2](left, right), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CompareNumbers[N1, N2](left, right)
		}
	})
}

func compareNumbersTestCaseName[N1, N2 numberType](left N1, right N2) string {
	return fmt.Sprintf("compare %[1]T(%[1]v) with %[2]T(%[2]v)", left, right)
}

type compareNumbersTestCase[N1, N2 numberType] struct {
	left     N1
	right    N2
	expected int
}
