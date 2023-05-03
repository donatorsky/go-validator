package validator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FieldsIterator(t *testing.T) {
	// given
	var (
		structData = someStruct{
			String: "foo",
			Slice:  []any{1, "foo", 2.0, true, nil},
			SliceOfSlices: [][]any{
				{1, "foo"},
				{2.0, true},
				{nil},
			},
			SliceOfMaps: []map[string]any{
				{"id": 1},
				{"id": 2},
				{"id": 3},
			},
			Array: [5]any{1, "foo", 2.0, true, nil},
		}
		mapData = map[string]any{
			"string":          structData.String,
			"slice":           structData.Slice,
			"slice_of_slices": structData.SliceOfSlices,
			"slice_of_maps":   structData.SliceOfMaps,
			"array":           structData.Array,
		}
	)

	// when
	for ttIdx, tt := range []struct {
		field          string
		data           any
		expectedValues []fieldValue
	}{
		{
			field: "non-existing",
			data:  mapData,
			expectedValues: []fieldValue{
				{
					pattern: false,
					field:   "non-existing",
					value:   nil,
					isNil:   true,
				},
			},
		},
		{
			field: "non-existing.*",
			data:  mapData,
			expectedValues: []fieldValue{
				{
					pattern: true,
					field:   "non-existing.*",
					value:   nil,
					isNil:   true,
				},
			},
		},
		{
			field: "non-existing.id",
			data:  mapData,
			expectedValues: []fieldValue{
				{
					pattern: false,
					field:   "non-existing.id",
					value:   nil,
					isNil:   true,
				},
			},
		},
		{
			field: "non-existing.*.id",
			data:  mapData,
			expectedValues: []fieldValue{
				{
					pattern: true,
					field:   "non-existing.*.id",
					value:   nil,
					isNil:   true,
				},
			},
		},
		{
			field: "non-existing.*.children.*",
			data:  mapData,
			expectedValues: []fieldValue{
				{
					pattern: true,
					field:   "non-existing.*.children.*",
					value:   nil,
					isNil:   true,
				},
			},
		},
		{
			field: "string",
			data:  nil,
			expectedValues: []fieldValue{
				{
					pattern: false,
					field:   "string",
					value:   nil,
					isNil:   true,
				},
			},
		},
		{
			field: "slice.id",
			data:  mapData,
			expectedValues: []fieldValue{
				{
					pattern: false,
					field:   "slice.id",
					value:   nil,
					isNil:   true,
				},
			},
		},
	} {
		t.Run(fmt.Sprintf("Non-existing test data #%d: %s from %T", ttIdx, tt.field, tt.data), func(t *testing.T) {
			var values []fieldValue

			for value := range newFieldsIterator(tt.field, tt.data) {
				values = append(values, value)
			}

			// then
			require.Equal(t, tt.expectedValues, values)
		})
	}

	// and when
	for _, data := range []any{
		mapData,
		structData,
	} {
		for ttIdx, tt := range []struct {
			field          string
			data           any
			expectedValues []fieldValue
		}{
			//{
			//	field: "*",
			//	data:  structData.Slice,
			//	expectedValues: []fieldValue{
			//		{
			//			pattern: true,
			//			field:   "0",
			//			value:   1,
			//			isNil:   false,
			//		},
			//		{
			//			pattern: true,
			//			field:   "1",
			//			value:   "foo",
			//			isNil:   false,
			//		},
			//		{
			//			pattern: true,
			//			field:   "2",
			//			value:   2.0,
			//			isNil:   false,
			//		},
			//		{
			//			pattern: true,
			//			field:   "3",
			//			value:   true,
			//			isNil:   false,
			//		},
			//		{
			//			pattern: true,
			//			field:   "4",
			//			value:   nil,
			//			isNil:   true,
			//		},
			//	},
			//},

			{
				field: "string",
				data:  data,
				expectedValues: []fieldValue{
					{
						pattern: false,
						field:   "string",
						value:   "foo",
						isNil:   false,
					},
				},
			},

			{
				field: "slice",
				data:  data,
				expectedValues: []fieldValue{
					{
						pattern: false,
						field:   "slice",
						value:   mapData["slice"],
						isNil:   false,
					},
				},
			},
			{
				field: "slice.*",
				data:  data,
				expectedValues: []fieldValue{
					{
						pattern: true,
						field:   "slice.0",
						value:   1,
						isNil:   false,
					},
					{
						pattern: true,
						field:   "slice.1",
						value:   "foo",
						isNil:   false,
					},
					{
						pattern: true,
						field:   "slice.2",
						value:   2.0,
						isNil:   false,
					},
					{
						pattern: true,
						field:   "slice.3",
						value:   true,
						isNil:   false,
					},
					{
						pattern: true,
						field:   "slice.4",
						value:   nil,
						isNil:   true,
					},
				},
			},
			{
				field: "slice_of_slices.*.*",
				data:  data,
				expectedValues: []fieldValue{
					{
						pattern: true,
						field:   "slice_of_slices.0.0",
						value:   1,
						isNil:   false,
					},
					{
						pattern: true,
						field:   "slice_of_slices.0.1",
						value:   "foo",
						isNil:   false,
					},
					{
						pattern: true,
						field:   "slice_of_slices.1.0",
						value:   2.0,
						isNil:   false,
					},
					{
						pattern: true,
						field:   "slice_of_slices.1.1",
						value:   true,
						isNil:   false,
					},
					{
						pattern: true,
						field:   "slice_of_slices.2.0",
						value:   nil,
						isNil:   true,
					},
				},
			},
			{
				field: "slice_of_maps.*.id",
				data:  data,
				expectedValues: []fieldValue{
					{
						pattern: true,
						field:   "slice_of_maps.0.id",
						value:   1,
						isNil:   false,
					},
					{
						pattern: true,
						field:   "slice_of_maps.1.id",
						value:   2,
						isNil:   false,
					},
					{
						pattern: true,
						field:   "slice_of_maps.2.id",
						value:   3,
						isNil:   false,
					},
				},
			},
			{
				field: "array",
				data:  data,
				expectedValues: []fieldValue{
					{
						pattern: false,
						field:   "array",
						value:   mapData["array"],
						isNil:   false,
					},
				},
			},
			{
				field: "array.*",
				data:  data,
				expectedValues: []fieldValue{
					{
						pattern: true,
						field:   "array.0",
						value:   1,
						isNil:   false,
					},
					{
						pattern: true,
						field:   "array.1",
						value:   "foo",
						isNil:   false,
					},
					{
						pattern: true,
						field:   "array.2",
						value:   2.0,
						isNil:   false,
					},
					{
						pattern: true,
						field:   "array.3",
						value:   true,
						isNil:   false,
					},
					{
						pattern: true,
						field:   "array.4",
						value:   nil,
						isNil:   true,
					},
				},
			},
		} {
			t.Run(fmt.Sprintf("Test data #%d: %s from %T", ttIdx, tt.field, tt.data), func(t *testing.T) {
				// when
				var values []fieldValue

				for value := range newFieldsIterator(tt.field, tt.data) {
					values = append(values, value)
				}

				// then
				require.Equal(t, tt.expectedValues, values)
			})
		}
	}
}

type someStruct struct {
	String        string           `validation:"string"`
	Slice         []any            `validation:"slice"`
	SliceOfSlices [][]any          `validation:"slice_of_slices"`
	SliceOfMaps   []map[string]any `validation:"slice_of_maps"`
	Array         [5]any           `validation:"array"`
}
