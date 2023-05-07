package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FieldsIterator_InvalidValue(t *testing.T) {
	// given
	for ttIdx, tt := range invalidValueTestCaseDataProvider() {
		t.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(t *testing.T) {
			var values []fieldValue

			// when
			for value := range newFieldsIterator(tt.field, tt.data) {
				values = append(values, value)
			}

			// then
			require.Equal(t, tt.expectedValues, values)
		})
	}
}

func Test_FieldsIterator_Slice(t *testing.T) {
	// given
	for ttIdx, tt := range sliceTestCaseDataProvider() {
		t.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(t *testing.T) {
			var values []fieldValue

			// when
			for value := range newFieldsIterator(tt.field, tt.data) {
				values = append(values, value)
			}

			// then
			require.Equal(t, tt.expectedValues, values)
		})
	}
}

func Test_FieldsIterator_Array(t *testing.T) {
	// given
	for ttIdx, tt := range arrayTestCaseDataProvider() {
		t.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(t *testing.T) {
			var values []fieldValue

			// when
			for value := range newFieldsIterator(tt.field, tt.data) {
				values = append(values, value)
			}

			// then
			require.Equal(t, tt.expectedValues, values)
		})
	}
}

func Test_FieldsIterator_Map(t *testing.T) {
	// given
	for ttIdx, tt := range mapTestCaseDataProvider() {
		t.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(t *testing.T) {
			var values []fieldValue

			// when
			for value := range newFieldsIterator(tt.field, tt.data) {
				values = append(values, value)
			}

			// then
			require.Equal(t, tt.expectedValues, values)
		})
	}
}

func Test_FieldsIterator_Struct(t *testing.T) {
	// given
	for ttIdx, tt := range structTestCaseDataProvider() {
		t.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(t *testing.T) {
			var values []fieldValue

			// when
			for value := range newFieldsIterator(tt.field, tt.data) {
				values = append(values, value)
			}

			// then
			require.Equal(t, tt.expectedValues, values)
		})
	}
}

func Test_FieldsIterator_Mixed(t *testing.T) {
	// given
	for ttIdx, tt := range mixedTestCaseDataProvider() {
		t.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(t *testing.T) {
			var values []fieldValue

			// when
			for value := range newFieldsIterator(tt.field, tt.data) {
				values = append(values, value)
			}

			// then
			require.Equal(t, tt.expectedValues, values)
		})
	}
}

func BenchmarkFieldsIterator_InvalidValue(b *testing.B) {
	for ttIdx, tt := range invalidValueTestCaseDataProvider() {
		b.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var values []fieldValue

				for value := range newFieldsIterator(tt.field, tt.data) {
					values = append(values, value)
				}
			}
		})
	}
}

func BenchmarkFieldsIterator_Slice(b *testing.B) {
	for ttIdx, tt := range sliceTestCaseDataProvider() {
		b.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var values []fieldValue

				for value := range newFieldsIterator(tt.field, tt.data) {
					values = append(values, value)
				}
			}
		})
	}
}

func BenchmarkFieldsIterator_Array(b *testing.B) {
	for ttIdx, tt := range arrayTestCaseDataProvider() {
		b.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var values []fieldValue

				for value := range newFieldsIterator(tt.field, tt.data) {
					values = append(values, value)
				}
			}
		})
	}
}

func BenchmarkFieldsIterator_Map(b *testing.B) {
	for ttIdx, tt := range mapTestCaseDataProvider() {
		b.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var values []fieldValue

				for value := range newFieldsIterator(tt.field, tt.data) {
					values = append(values, value)
				}
			}
		})
	}
}

func BenchmarkFieldsIterator_Struct(b *testing.B) {
	for ttIdx, tt := range structTestCaseDataProvider() {
		b.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var values []fieldValue

				for value := range newFieldsIterator(tt.field, tt.data) {
					values = append(values, value)
				}
			}
		})
	}
}

func BenchmarkFieldsIterator_Mixed(b *testing.B) {
	for ttIdx, tt := range mixedTestCaseDataProvider() {
		b.Run(fmt.Sprintf("#%d: %s on %s", ttIdx, tt.field, getType(tt.data)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var values []fieldValue

				for value := range newFieldsIterator(tt.field, tt.data) {
					values = append(values, value)
				}
			}
		})
	}
}

type testCaseData struct {
	field          string
	data           any
	expectedValues []fieldValue
}

func invalidValueTestCaseDataProvider() (tcd []testCaseData) {
	for _, field := range []string{
		"0",
		"*",
		"*.0",
		"*.*",

		"non-existing",
		"non-existing.0",
		"non-existing.bar",
		"non-existing.*",

		"non-existing.0.0",
		"non-existing.0.baz",
		"non-existing.0.*",

		"non-existing.bar.0",
		"non-existing.bar.baz",
		"non-existing.bar.*",

		"non-existing.*.0",
		"non-existing.*.baz",
		"non-existing.*.*",

		"non-existing.*.0.*",
		"non-existing.*.bar.*",
	} {
		for _, data := range []any{
			nil,
			"some string",
			ptr("some string"),
		} {
			tcd = append(tcd, testCaseData{
				field: field,
				data:  data,
				expectedValues: []fieldValue{
					{
						field: field,
						value: nil,
					},
				},
			})
		}
	}

	return
}

func sliceTestCaseDataProvider() []testCaseData {
	return []testCaseData{
		{
			field: "foo",
			data:  ([]any)(nil),
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "foo",
			data:  []any{"foo"},
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "foo",
			data:  ptr([]any{"foo"}),
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "0",
			data:  []any{"foo"},
			expectedValues: []fieldValue{
				{
					field: "0",
					value: "foo",
				},
			},
		},
		{
			field: "0",
			data:  ptr([]any{"foo"}),
			expectedValues: []fieldValue{
				{
					field: "0",
					value: "foo",
				},
			},
		},
		{
			field: "1",
			data:  []any{"foo"},
			expectedValues: []fieldValue{
				{
					field: "1",
					value: nil,
				},
			},
		},
		{
			field: "1",
			data:  ptr([]any{"foo"}),
			expectedValues: []fieldValue{
				{
					field: "1",
					value: nil,
				},
			},
		},
		{
			field: "*",
			data:  []any{"foo", "bar"},
			expectedValues: []fieldValue{
				{
					field: "0",
					value: "foo",
				},
				{
					field: "1",
					value: "bar",
				},
			},
		},
		{
			field: "*",
			data:  ptr([]any{"foo", "bar"}),
			expectedValues: []fieldValue{
				{
					field: "0",
					value: "foo",
				},
				{
					field: "1",
					value: "bar",
				},
			},
		},
		{
			field: "*.0",
			data: []any{
				[]string{"foo1"},
				[]string{"foo2"},
			},
			expectedValues: []fieldValue{
				{
					field: "0.0",
					value: "foo1",
				},
				{
					field: "1.0",
					value: "foo2",
				},
			},
		},
		{
			field: "*.foo",
			data: []any{
				[]string{"foo1"},
				[]string{"foo2"},
			},
			expectedValues: []fieldValue{
				{
					field: "0.foo",
					value: nil,
				},
				{
					field: "1.foo",
					value: nil,
				},
			},
		},
		{
			field: "*.*",
			data: []any{
				[]string{"foo1", "bar1"},
				[]string{"foo2", "bar2"},
			},
			expectedValues: []fieldValue{
				{
					field: "0.0",
					value: "foo1",
				},
				{
					field: "0.1",
					value: "bar1",
				},
				{
					field: "1.0",
					value: "foo2",
				},
				{
					field: "1.1",
					value: "bar2",
				},
			},
		},
		{
			field: "*.*",
			data: ptr([]any{
				ptr([]*string{ptr("foo1"), ptr("bar1")}),
				ptr([]*string{ptr("foo2"), ptr("bar2")}),
			}),
			expectedValues: []fieldValue{
				{
					field: "0.0",
					value: "foo1",
				},
				{
					field: "0.1",
					value: "bar1",
				},
				{
					field: "1.0",
					value: "foo2",
				},
				{
					field: "1.1",
					value: "bar2",
				},
			},
		},
		{
			field: "*.*.foo",
			data: []any{
				[]string{"foo1", "bar1"},
				[]string{"foo2", "bar2"},
			},
			expectedValues: []fieldValue{
				{
					field: "0.0.foo",
					value: nil,
				},
				{
					field: "0.1.foo",
					value: nil,
				},
				{
					field: "1.0.foo",
					value: nil,
				},
				{
					field: "1.1.foo",
					value: nil,
				},
			},
		},
	}
}

func arrayTestCaseDataProvider() []testCaseData {
	return []testCaseData{
		{
			field: "foo",
			data:  [0]any{},
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "foo",
			data:  ptr([0]any{}),
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "foo",
			data:  [1]any{"foo"},
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "foo",
			data:  ptr([1]any{"foo"}),
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "0",
			data:  [1]any{"foo"},
			expectedValues: []fieldValue{
				{
					field: "0",
					value: "foo",
				},
			},
		},
		{
			field: "0",
			data:  ptr([1]any{"foo"}),
			expectedValues: []fieldValue{
				{
					field: "0",
					value: "foo",
				},
			},
		},
		{
			field: "1",
			data:  [1]any{"foo"},
			expectedValues: []fieldValue{
				{
					field: "1",
					value: nil,
				},
			},
		},
		{
			field: "1",
			data:  ptr([1]any{"foo"}),
			expectedValues: []fieldValue{
				{
					field: "1",
					value: nil,
				},
			},
		},
		{
			field: "*",
			data:  [2]any{"foo", "bar"},
			expectedValues: []fieldValue{
				{
					field: "0",
					value: "foo",
				},
				{
					field: "1",
					value: "bar",
				},
			},
		},
		{
			field: "*",
			data:  ptr([2]any{"foo", "bar"}),
			expectedValues: []fieldValue{
				{
					field: "0",
					value: "foo",
				},
				{
					field: "1",
					value: "bar",
				},
			},
		},
		{
			field: "*.0",
			data: [2]any{
				[1]string{"foo1"},
				[1]string{"foo2"},
			},
			expectedValues: []fieldValue{
				{
					field: "0.0",
					value: "foo1",
				},
				{
					field: "1.0",
					value: "foo2",
				},
			},
		},
		{
			field: "*.foo",
			data: [2]any{
				[1]string{"foo1"},
				[1]string{"foo2"},
			},
			expectedValues: []fieldValue{
				{
					field: "0.foo",
					value: nil,
				},
				{
					field: "1.foo",
					value: nil,
				},
			},
		},
		{
			field: "*.*",
			data: [2]any{
				[2]string{"foo1", "bar1"},
				[2]string{"foo2", "bar2"},
			},
			expectedValues: []fieldValue{
				{
					field: "0.0",
					value: "foo1",
				},
				{
					field: "0.1",
					value: "bar1",
				},
				{
					field: "1.0",
					value: "foo2",
				},
				{
					field: "1.1",
					value: "bar2",
				},
			},
		},
		{
			field: "*.*",
			data: ptr([2]any{
				ptr([2]*string{ptr("foo1"), ptr("bar1")}),
				ptr([2]*string{ptr("foo2"), ptr("bar2")}),
			}),
			expectedValues: []fieldValue{
				{
					field: "0.0",
					value: "foo1",
				},
				{
					field: "0.1",
					value: "bar1",
				},
				{
					field: "1.0",
					value: "foo2",
				},
				{
					field: "1.1",
					value: "bar2",
				},
			},
		},
		{
			field: "*.*.foo",
			data: [2]any{
				[2]string{"foo1", "bar1"},
				[2]string{"foo2", "bar2"},
			},
			expectedValues: []fieldValue{
				{
					field: "0.0.foo",
					value: nil,
				},
				{
					field: "0.1.foo",
					value: nil,
				},
				{
					field: "1.0.foo",
					value: nil,
				},
				{
					field: "1.1.foo",
					value: nil,
				},
			},
		},
	}
}

func mapTestCaseDataProvider() []testCaseData {
	return []testCaseData{
		{
			field: "foo",
			data:  (map[string]any)(nil),
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "foo",
			data:  map[string]any{},
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "foo",
			data: map[string]any{
				"bar": "baz",
			},
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "foo",
			data: ptr(map[string]any{
				"bar": "baz",
			}),
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "foo",
			data: map[string]any{
				"foo": "bar",
			},
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: "bar",
				},
			},
		},
		{
			field: "foo",
			data: ptr(map[string]any{
				"foo": "bar",
			}),
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: "bar",
				},
			},
		},
		{
			field: "0",
			data: map[string]any{
				"foo": "bar",
			},
			expectedValues: []fieldValue{
				{
					field: "0",
					value: nil,
				},
			},
		},
		{
			field: "0",
			data: ptr(map[string]any{
				"foo": "bar",
			}),
			expectedValues: []fieldValue{
				{
					field: "0",
					value: nil,
				},
			},
		},
		{
			field: "*",
			data: map[string]any{
				"foo": "bar",
			},
			expectedValues: []fieldValue{
				{
					field: "*",
					value: nil,
				},
			},
		},
		{
			field: "*",
			data: ptr(map[string]any{
				"foo": "bar",
			}),
			expectedValues: []fieldValue{
				{
					field: "*",
					value: nil,
				},
			},
		},
		{
			field: "*.0",
			data: map[string]any{
				"foo": "bar",
			},
			expectedValues: []fieldValue{
				{
					field: "*.0",
					value: nil,
				},
			},
		},
		{
			field: "*.foo",
			data: map[string]any{
				"foo": "bar",
			},
			expectedValues: []fieldValue{
				{
					field: "*.foo",
					value: nil,
				},
			},
		},
		{
			field: "*.*",
			data: map[string]any{
				"foo": "bar",
			},
			expectedValues: []fieldValue{
				{
					field: "*.*",
					value: nil,
				},
			},
		},
		{
			field: "foo.bar",
			data: map[string]any{
				"foo": map[string]any{
					"bar": "baz",
				},
			},
			expectedValues: []fieldValue{
				{
					field: "foo.bar",
					value: "baz",
				},
			},
		},
		{
			field: "foo.*",
			data: map[string]any{
				"foo": map[string]any{
					"bar": "baz",
				},
			},
			expectedValues: []fieldValue{
				{
					field: "foo.*",
					value: nil,
				},
			},
		},
		{
			field: "foo.bar.baz",
			data: map[string]any{
				"foo": map[string]any{
					"bar": map[string]any{
						"baz": "lorem",
					},
				},
			},
			expectedValues: []fieldValue{
				{
					field: "foo.bar.baz",
					value: "lorem",
				},
			},
		},
		{
			field: "foo.bar.baz",
			data: ptr(map[string]any{
				"foo": ptr(map[string]any{
					"bar": ptr(map[string]any{
						"baz": "lorem",
					}),
				}),
			}),
			expectedValues: []fieldValue{
				{
					field: "foo.bar.baz",
					value: "lorem",
				},
			},
		},
		{
			field: "foo.bar.*",
			data: map[string]any{
				"foo": map[string]any{
					"bar": map[string]any{
						"baz": "lorem",
					},
				},
			},
			expectedValues: []fieldValue{
				{
					field: "foo.bar.*",
					value: nil,
				},
			},
		},
	}
}

func structTestCaseDataProvider() []testCaseData {
	return []testCaseData{
		{
			field: "foo",
			data:  (*someStruct)(nil),
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: nil,
				},
			},
		},
		{
			field: "Foo",
			data:  (*someStruct)(nil),
			expectedValues: []fieldValue{
				{
					field: "Foo",
					value: nil,
				},
			},
		},
		{
			field: "foo",
			data: someStruct{
				FooNamed: "foo",
			},
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: "foo",
				},
			},
		},
		{
			field: "Foo",
			data: someStruct{
				Foo: "Foo",
			},
			expectedValues: []fieldValue{
				{
					field: "Foo",
					value: "Foo",
				},
			},
		},
		{
			field: "foo",
			data: &someStruct{
				FooNamed: "foo",
			},
			expectedValues: []fieldValue{
				{
					field: "foo",
					value: "foo",
				},
			},
		},
		{
			field: "Foo",
			data: &someStruct{
				Foo: "Foo",
			},
			expectedValues: []fieldValue{
				{
					field: "Foo",
					value: "Foo",
				},
			},
		},
		{
			field: "Struct",
			data: someStruct{
				Struct:      nil,
				StructNamed: &someStruct{},
			},
			expectedValues: []fieldValue{
				{
					field: "Struct",
					value: nil,
				},
			},
		},
		{
			field: "Struct.Foo",
			data: someStruct{
				Struct: nil,
				StructNamed: &someStruct{
					Foo: "Foo",
				},
			},
			expectedValues: []fieldValue{
				{
					field: "Struct.Foo",
					value: nil,
				},
			},
		},
		{
			field: "struct",
			data: someStruct{
				Struct: nil,
				StructNamed: &someStruct{
					Foo: "foo",
				},
			},
			expectedValues: []fieldValue{
				{
					field: "struct",
					value: someStruct{
						Foo: "foo",
					},
				},
			},
		},
		{
			field: "struct.Foo",
			data: someStruct{
				Struct: nil,
				StructNamed: &someStruct{
					Foo: "Foo",
				},
			},
			expectedValues: []fieldValue{
				{
					field: "struct.Foo",
					value: "Foo",
				},
			},
		},
		{
			field: "struct.Struct.foo",
			data: someStruct{
				Struct: nil,
				StructNamed: &someStruct{
					Struct: &someStruct{
						FooNamed: "foo",
					},
				},
			},
			expectedValues: []fieldValue{
				{
					field: "struct.Struct.foo",
					value: "foo",
				},
			},
		},
		{
			field: "any.Struct.foo",
			data: someStruct{
				Any: someStruct{
					Struct: &someStruct{
						FooNamed: "foo",
					},
				},
			},
			expectedValues: []fieldValue{
				{
					field: "any.Struct.foo",
					value: "foo",
				},
			},
		},
	}
}

func mixedTestCaseDataProvider() []testCaseData {
	return []testCaseData{
		// slice
		{
			field: "*.*",
			data: [][1]int{
				{1},
				{2},
			},
			expectedValues: []fieldValue{
				{
					field: "0.0",
					value: 1,
				},
				{
					field: "1.0",
					value: 2,
				},
			},
		},
		{
			field: "*.*",
			data: ptr([]*[1]*int{
				{ptr(1)},
				{ptr(2)},
			}),
			expectedValues: []fieldValue{
				{
					field: "0.0",
					value: 1,
				},
				{
					field: "1.0",
					value: 2,
				},
			},
		},
		{
			field: "*.0",
			data: [][1]int{
				{1},
				{2},
			},
			expectedValues: []fieldValue{
				{
					field: "0.0",
					value: 1,
				},
				{
					field: "1.0",
					value: 2,
				},
			},
		},
		{
			field: "*.foo",
			data: []someStruct{
				{
					Foo:      "foo_1",
					FooNamed: "foo_named_1",
				},
				{
					Foo:      "foo_2",
					FooNamed: "foo_named_2",
				},
			},
			expectedValues: []fieldValue{
				{
					field: "0.foo",
					value: "foo_named_1",
				},
				{
					field: "1.foo",
					value: "foo_named_2",
				},
			},
		},
		{
			field: "*.foo",
			data: []*someStruct{
				{
					Foo:      "foo_1",
					FooNamed: "foo_named_1",
				},
				{
					Foo:      "foo_2",
					FooNamed: "foo_named_2",
				},
			},
			expectedValues: []fieldValue{
				{
					field: "0.foo",
					value: "foo_named_1",
				},
				{
					field: "1.foo",
					value: "foo_named_2",
				},
			},
		},
		{
			field: "*.Foo",
			data: []someStruct{
				{
					Foo:      "foo_1",
					FooNamed: "foo_named_1",
				},
				{
					Foo:      "foo_2",
					FooNamed: "foo_named_2",
				},
			},
			expectedValues: []fieldValue{
				{
					field: "0.Foo",
					value: "foo_1",
				},
				{
					field: "1.Foo",
					value: "foo_2",
				},
			},
		},
		{
			field: "*.foo",
			data: []map[string]int{
				{
					"foo": 1,
				},
				{
					"foo": 2,
				},
			},
			expectedValues: []fieldValue{
				{
					field: "0.foo",
					value: 1,
				},
				{
					field: "1.foo",
					value: 2,
				},
			},
		},

		// array
		{
			field: "*.*",
			data: [2][]int{
				{1},
				{2},
			},
			expectedValues: []fieldValue{
				{
					field: "0.0",
					value: 1,
				},
				{
					field: "1.0",
					value: 2,
				},
			},
		},
		{
			field: "*.*",
			data: ptr([2]*[]*int{
				{ptr(1)},
				{ptr(2)},
			}),
			expectedValues: []fieldValue{
				{
					field: "0.0",
					value: 1,
				},
				{
					field: "1.0",
					value: 2,
				},
			},
		},

		// map
		{
			field: "foo.*",
			data: map[string][]int{
				"foo": {1, 2},
				"bar": {3},
			},
			expectedValues: []fieldValue{
				{
					field: "foo.0",
					value: 1,
				},
				{
					field: "foo.1",
					value: 2,
				},
			},
		},
		{
			field: "foo.*",
			data: ptr(map[string]*[]*int{
				"foo": {ptr(1), ptr(2)},
				"bar": {ptr(3)},
			}),
			expectedValues: []fieldValue{
				{
					field: "foo.0",
					value: 1,
				},
				{
					field: "foo.1",
					value: 2,
				},
			},
		},
		{
			field: "foo.Foo",
			data: map[string]someStruct{
				"foo": {
					Foo: "baz1",
				},
				"bar": {
					Foo: "baz2",
				},
			},
			expectedValues: []fieldValue{
				{
					field: "foo.Foo",
					value: "baz1",
				},
			},
		},
		{
			field: "foo.*.Foo",
			data: map[string][]*someStruct{
				"foo": {
					{
						Foo: "bar",
					},
					{
						Foo: "baz",
					},
				},
				"bar": nil,
			},
			expectedValues: []fieldValue{
				{
					field: "foo.0.Foo",
					value: "bar",
				},
				{
					field: "foo.1.Foo",
					value: "baz",
				},
			},
		},

		// struct
		{
			field: "any.*",
			data: someStruct{
				Any: []int{1, 2},
			},
			expectedValues: []fieldValue{
				{
					field: "any.0",
					value: 1,
				},
				{
					field: "any.1",
					value: 2,
				},
			},
		},
		{
			field: "any.foo",
			data: someStruct{
				Any: map[string]int{
					"foo": 1,
				},
			},
			expectedValues: []fieldValue{
				{
					field: "any.foo",
					value: 1,
				},
			},
		},
		{
			field: "any.*.foo",
			data: someStruct{
				Any: []*someStruct{
					{
						FooNamed: "bar",
					},
					{
						FooNamed: "baz",
					},
				},
			},
			expectedValues: []fieldValue{
				{
					field: "any.0.foo",
					value: "bar",
				},
				{
					field: "any.1.foo",
					value: "baz",
				},
			},
		},
	}
}

type someStruct struct {
	Foo         string
	FooNamed    string `validation:"foo"`
	Struct      *someStruct
	StructNamed *someStruct `validation:"struct"`
	Any         any         `validation:"any"`
}

func getType(v any) string {
	valueOf := reflect.ValueOf(v)

	if !valueOf.IsValid() {
		return "nil"
	}

	for valueOf.Kind() == reflect.Pointer {
		if valueOf.IsNil() {
			return valueOf.Type().String()
		} else {
			return "*" + getType(valueOf.Elem().Interface())
		}
	}

	switch typeOf := reflect.TypeOf(v); typeOf.Kind() {
	case reflect.Slice, reflect.Array:
		var size string
		var elements []string

		if valueOf.IsZero() {
			size = "nil"
		} else {
			size = strconv.Itoa(valueOf.Len())

			elements = make([]string, valueOf.Len())
			for idx := 0; idx < valueOf.Len(); idx++ {
				elements[idx] = getType(valueOf.Index(idx).Interface())
			}
		}

		typeName := typeOf.Elem().String()
		if typeName == "interface {}" {
			typeName = "any"
		} else {
			typeName = strings.Replace(typeName, "validator.", "", 1)
		}

		return fmt.Sprintf(
			"[%s]%s{%s}",
			size,
			typeName,
			strings.Join(elements, ","),
		)

	case reflect.Map:
		var elements []string

		if !valueOf.IsZero() {
			elements = make([]string, 0, valueOf.Len())
			for _, key := range valueOf.MapKeys() {
				elements = append(elements, fmt.Sprintf("%q:%s", key.String(), getType(valueOf.MapIndex(key).Interface())))
			}
		}

		keyName := typeOf.Key().String()
		if keyName == "interface {}" {
			keyName = "any"
		} else {
			keyName = strings.Replace(keyName, "validator.", "", 1)
		}

		typeName := typeOf.Elem().String()
		if typeName == "interface {}" {
			typeName = "any"
		} else {
			typeName = strings.Replace(typeName, "validator.", "", 1)
		}

		return fmt.Sprintf(
			"[%s]%s{%s}",
			keyName,
			typeName,
			strings.Join(elements, ","),
		)

	default:
		return strings.Replace(typeOf.String(), "validator.", "", 1)
	}
}

func ptr[T any](v T) *T {
	return &v
}
