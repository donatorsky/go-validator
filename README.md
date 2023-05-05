# Go Validator
A Laravel-like data validator for Go.

[![GitHub license](https://img.shields.io/github/license/donatorsky/go-validator)](https://github.com/donatorsky/go-validator/blob/main/LICENSE)
[![Build](https://github.com/donatorsky/go-validator/workflows/Tests/badge.svg?branch=main)](https://github.com/donatorsky/go-validator/actions?query=branch%3Amain)

Allows you to define a list of validation rules (constraints) for a specific field and produces a summary with failures.

A rule can change the value it validates (only during validation, the original value remains unchanged), which can be beneficial in later validation.
Some rules adapt to the currently validated value and act differently. More details can be found in descriptions of rules.

## Installation

```shell
go get github.com/donatorsky/go-validator
```

The library is still a work in progress. More details soon.

## Available validators

Each validator has a context counterpart which lets you set the context used during validation. It will be passed to rules. The default context is `context.Background()`.

### `ForMap`, `ForMapWithContext`

Validates a `map[string]any`. You can specify a map of rules for each key and internal values (either slices, arrays, maps or structs).

Returns `ErrorsBag` with keys being the map keys of input map.

#### Example
```go
validator.ForMap(
    map[string]any{
        "foo": 123,
        "bar": "bar",
        "baz": map[string]any{
            "inner_foo": 123,
            "inner_bar": "bar",
            // ...
        },
    },
    validator.RulesMap{
        "foo":   {
            rule.Required(),
            rule.Integer[int](),
        },

        "bar":   {
            rule.Required(),
            rule.String(),
        },

        "baz.inner_foo":   {
            rule.Required(),
            rule.Integer[int](),
        },

        "baz.inner_bar":   {
            rule.Required(),
            rule.Slice(),
        },

        // ...
    },
)
```

### `ForStruct`, `ForStructWithContext`

Validates a struct. You can specify a map of rules for each field name and internal values (either slices, arrays, maps or structs).

Note that you can also use custom name for a field using `validation` tag. 

You can also pass pointer which will be automatically dereferenced.

Returns `ErrorsBag` with keys being the field names (or values from `validation` if provided) of input struct.

#### Example
```go
type SomeRequest struct {
    Foo int
    Bar string `validation:"bar"`
    Baz SomeRequestBaz
}

type SomeRequestBaz struct {
    InnerFoo int `validation:"foo"`
    InnerBar string
    // ...
}

validator.ForStruct(
    SomeRequest{
        Foo: 123,
        Bar: "bar",
        Baz: SomeRequestBaz{
            InnerFoo: 123,
            InnerBar: "bar",
            // ...
        },
    },
    validator.RulesMap{
        "Foo":   {
            rule.Required(),
            rule.Integer[int](),
        },

        "bar":   {
            rule.Required(),
            rule.String(),
        },

        "Baz.foo":   {
            rule.Required(),
            rule.Integer[int](),
        },

        "Baz.InnerBar":   {
            rule.Required(),
            rule.Slice(),
        },

        // ...
    },
)
```

### `ForSlice`, `ForSliceWithContext`

Validates both a slice `[]any` or an array `[size]any`. You can specify a list of rules for each element in given slice/array.

You can also pass pointer which will be automatically dereferenced.

Returns `ErrorsBag` with keys being the indices of input slice/array.

#### Example
```go
validator.ForSlice(
    []any{
        "foo",
        "bar",
    },
    validator.RulesMap{
        rule.Required(),
        rule.String(),
        // ...
    },
)
```

### `ForValue`, `ForValueWithContext`

Validates any value. You can specify a list of rules for given value.

If you pass a pointer, it will not be dereferenced.

Returns a slice of `ValidationError`.

#### Example
```go
validator.ForValue(
    123,
    validator.RulesMap{
        rule.Required(),
        rule.Integer[int](),
        rule.Min(100),
        // ...
    },
)
```

## Validation of nested objects

When map or struct contains a nested object (e.g.: slice, array, map or struct), you can also validate every single value of it by using `*` wildcard symbol.

#### Example
```go
type SomeRequest struct {
    SingleValue    int
    Array          [3]string
    Slice          []string
    Map            map[string]string
    Struct         SomeRequestStruct
    SliceOfStructs []SomeRequestStruct
    MapOfStructs   map[string]SomeRequestStruct
    SliceOfSlices  [][]string
}

type SomeRequestStruct struct {
    InnerSingleValue int
    InnerArray       [3]string
    InnerSlice       []string
    InnerMap         map[string]string
    InnerStruct      SomeRequestInnerStruct
}

type SomeRequestInnerStruct struct {
    InnerInnerSingleValue int
    InnerInnerArray       [3]string
    InnerInnerSlice       []string
    InnerInnerMap         map[string]string
    // ...
}

validator.ForStruct(
    SomeRequest{
        SingleValue: 123,
        Array:       [3]string{"foo", "bar", "baz"},
        Slice:       []string{"foo", "bar", "baz"},
        Map:         map[string]string{"foo": 1, "bar": 2},
        Struct: SomeRequestStruct{
            InnerSingleValue: 123,
            InnerArray:       [3]string{"foo", "bar", "baz"},
            InnerSlice:       []string{"foo", "bar", "baz"},
            InnerMap:         map[string]string{"foo": 1, "bar": 2},
            InnerStruct:      SomeRequestInnerStruct{
                // ...
            },
        },
        SliceOfStructs: []SomeRequestStruct{
            {
                InnerSingleValue: 123,
                InnerSlice:       []string{"foo", "bar", "baz"},
                // ...
            },
            // ...
        },
        MapOfStructs: map[string]SomeRequestStruct{
            "foo": {
                InnerSingleValue: 123,
                InnerSlice:       []string{"foo", "bar", "baz"},
                // ...
            },
            "bar": {
                InnerSingleValue: 123,
                InnerSlice:       []string{"foo", "bar", "baz"},
                // ...
            },
            // ...
        },
        SliceOfSlices: [][]string{
            {"foo", "bar", "baz"},
            {"foo", "bar"},
        },
    },
    validator.RulesMap{
        // Rules for SomeRequest
        "SingleValue": {
            rule.Required(),
            rule.Integer[int](),
        },

        "Array": {
            rule.Required(),
            rule.SliceOf[string](),
        },

        "Array.*": {
            rule.Required(),
            rule.String(),
        },

        "Slice": {
            rule.Required(),
            rule.SliceOf[string](),
            rule.Length(3),
        },

        "Slice.*": {
            rule.Required(),
            rule.String(),
        },

        "Map.foo": {
            rule.Required(),
            rule.Integer[int](),
        },

        "Map.bar": {
            rule.Required(),
            rule.Integer[int](),
        },

        // Rules for SomeRequestStruct
        "Struct.InnerSingleValue": {
            rule.Required(),
            rule.Integer[int](),
        },

        "Struct.InnerArray": {
            rule.Required(),
            rule.SliceOf[string](),
        },

        "Struct.InnerArray.*": {
            rule.Required(),
            rule.String(),
        },

        "Struct.InnerSlice": {
            rule.Required(),
            rule.SliceOf[string](),
            rule.Length(3),
        },

        "Struct.InnerSlice.*": {
            rule.Required(),
            rule.String(),
        },

        "Struct.InnerMap.foo": {
            rule.Required(),
            rule.Integer[int](),
        },

        "Struct.InnerMap.bar": {
            rule.Required(),
            rule.Integer[int](),
        },

        // Rules for SomeRequest (complex examples)
        "SliceOfStructs": {
            rule.Required(),
            rule.Length(1),
        },

        "SliceOfStructs.*.InnerSingleValue": {
            rule.Required(),
            rule.Integer[int](),
        },

        "SliceOfStructs.*.InnerSlice.*": {
            rule.Required(),
            rule.String(),
        },

        "MapOfStructs": {
            rule.Required(),
        },

        "MapOfStructs.foo.InnerSingleValue": {
            rule.Required(),
            rule.Integer[int](),
        },

        "MapOfStructs.foo.InnerSlice.*": {
            rule.Required(),
            rule.String(),
        },

        "SliceOfSlices": {
            rule.Required(),
            rule.Slice(),
            rule.Length(2),
        },

        "SliceOfSlices.*": {
            rule.Required(),
            rule.Slice(),
            rule.Min(2),
        },

        "SliceOfSlices.*.*": {
            rule.Required(),
            rule.String(),
        },

        // Non-existing keys
        "IDoNotExist": {
            rule.Required(),
        },

        "IDoNotExist.*": {
            rule.Required(),
        },

        "IDoNotExist.foo": {
            rule.Required(),
        },

        "SliceOfSlices.*.*.*": {
            rule.Required(),
        },

        "SliceOfSlices.*.*.foo": {
            rule.Required(),
        },

        // ...
    },
)
```

As a result you will get errors for each element separately, e.g.: `SliceOfSlices.0.1`, `SliceOfSlices.3.0` etc.
Note that some wildcards will not be matched. In that case you will get `*` for every unmatched nested element, e.g.: `IDoNotExist.*`, `"IDoNotExist.foo"`, `"SliceOfSlices.0.0.*"`, `"SliceOfSlices.0.1.*"`, `"SliceOfSlices.0.0.foo"`, `"SliceOfSlices.0.1.foo"` etc.

## Stopping validation on first error

Some rules stop validation of given element once they fail (e.g.: `Required` since further validation makes no sense when value is not present).

You can also manually stop validation by using `Bail` pseudo-rule.

#### Example
```go
validator.ForValue(
    "123",
    validator.RulesMap{
        rule.Required(),
        rule.Integer[int](),
        rule.Bail(), // Next rules will not be checked if value is not an integer
        rule.Min(100),
        // ...
    },
)
```

## Conditional validation

You can add validation rules based on custom conditions. It can be either simple boolean value using `When` or complex condition using `WhenFunc`.

Conditional rules can be nested to cover more complex requirements.

Once conditional rule is valid, its rules will be merged to the main list of rules. It also includes the `Bail` pseudo-rule which will stop any further validation of given rules list, no matter how deep it was defined.

### `When`

This pseudo-rule allows for adding rules based on simple boolean value.

#### Example
```go
validator.ForValue(
    123,
    validator.RulesMap{
        rule.Required(),
        rule.Integer[int](),
        rule.When(true,
            rule.Min(100),
            rule.When(false,
                rule.Min(200),
                // ...
            ),
            rule.Bail(),
            // ...
        ),
        // ...
    },
)
```

The example above becomes:
```go
validator.ForValue(
    123,
    validator.RulesMap{
        rule.Required(),
        rule.Integer[int](),
        rule.Min(100),
        rule.Bail(),
        // ...
        // ...
    },
)
```

### `WhenFunc`

This pseudo-rule allows for adding rules based on a custom logic.

It receives the `context` passed to the validator, the currently validated `value` and the original `data` passed to the validator.

#### Example
```go
validator.ForValue(
    123,
    validator.RulesMap{
        rule.Required(),
        rule.Integer[int](),
        rule.WhenFunc[int](
            func(ctx context.Context, value int, data any) bool {
                return value%2 == 1
            },
            rule.Min(100),
            rule.WhenFunc[int](
                func(_ context.Context, value int, _ any) bool {
                    return value%2 == 0
                },
                rule.Min(200),
                // ...
            ),
            rule.Bail(),
            // ...
        ),
        // ...
    },
)
```

The example above becomes:
```go
validator.ForValue(
    123,
    validator.RulesMap{
        rule.Required(),
        rule.Integer[int](),
        rule.Min(100),
        rule.Bail(),
        // ...
        // ...
    },
)
```

## Available rules

TODO (see [rules](rule) directory).

## Custom validation

You can write a custom validator to cover custom needs. There are to ways of doing it: by implementing `rule.Rule` interface or by using `rule.Custom` rule.

A custom rule struct can also implement `BailingRule` interface so that it may stop further validation. There is also `Bailer` helper struct for that.

The `rule.Custom` rule can return any `error`. In that case, the error is added to the response. However, you can return a custom message by returning an error of `error.ValidationError` type.

Since the value can be anything, including pointer, there is a helper function `rule.Dereference` that returns the underlying value.

#### Example
```go
import ve "github.com/donatorsky/go-validator/error"

type DividesByNValidationError struct {
    ve.BasicValidationError

    Divider int `json:"divider"`
}

func (e DividesByNValidationError) Error() string {
    return fmt.Sprintf("Cannot be divided by %d", e.Divider)
}

type DividesByN struct {
    divider int
}

func (r DividesByN) Apply(_ context.Context, value any, _ any) (any, ve.ValidationError) {
    v, isNil := Dereference(value)
    if isNil {
        return value, nil
    }

    if v.(int) % r.divider != 0 {
        return value, &DividesByNValidationError{
            BasicValidationError: ve.BasicValidationError{
                Rule: ve.TypeCustom,
            },
            Divider: r.divider,
        }
    }
    
    return value, nil
}

validator.ForValue(
    123,
    validator.RulesMap{
        rule.Required(),
        rule.Integer[int](),
        rule.Custom(func(_ context.Context, value int, _ any) (newValue int, err error) {
            switch value % 2 {
            case 0:
                return value, nil

            case 1:
                return value + 1, nil
            }
        }),
        rule.Min(124), // passes because custom rule modified the value by adding 1
        &DividesByN{divider: 3}, // fails
        // ...
    },
)
```

## Example

```go
package main

import (
    "context"
    "errors"
    "encoding/json"
    "fmt"
    "time"

    "github.com/donatorsky/go-validator"
    "github.com/donatorsky/go-validator/rule"
)

func main() {
	var (
		mapData = map[string]any{
			"int":         int(125),
			"*int":        ptr(int(124)),
			"**int":       ptr(ptr(int(124))),
			"string":      "Lorem ipsum",
			"boolean":     true,
			"skip":        "123",
			"date_string": "2023-03-31",
			"array":       []any{"Foo", nil, "Bar", "Baz", 123},
			"child": map[string]any{
				"id": 123,
				"ancestor": map[string]int{
					"id": 456,
				},
			},
			"children": []map[string]any{
				{"id": 123},
				{"id": 456},
				{"id": 789},
			},
			"object": data{
				ID:      -1,
				Name:    "value: Name",
				Enabled: true,
				Skip:    "value: Skip",
				NoName:  "value: NoName",
			},
		}

		intRules = []rule.Rule{
			rule.When(true,
				rule.When(true,
					rule.Min(101),
					//rule.Bail(),
				),
				rule.Min(150),
				//rule.Bail(),
			),
			rule.Required(),
			rule.Integer[int](),
			//rule.Bail(),
			rule.Min(1240),
			rule.Min(float32(125.01)),
			rule.Min(float64(125.01)),
			rule.When(true,
				rule.Min(102),
				//rule.Bail(),
				rule.When(true,
					rule.Min(103),
					//rule.Bail(),
				),
			),
			rule.WhenFunc(
				func(_ context.Context, value int, _ any) bool {
					time.Sleep(time.Millisecond * 100)

					return value%2 == 0
				},
				rule.Min(200),
				//rule.Bail(),
				rule.When(true,
					rule.When(true,
						rule.Min(300),
						//rule.Bail(),
					),
					rule.Max(90),
					//rule.Bail(),
				),
				rule.Max(100),
			),
			rule.Min(400),
			rule.Custom(func(_ context.Context, value int, _ any) (int, error) {
				switch value % 3 {
				case 0:
					return value + 1, nil

				case 1:
					return value - 1, nil

				default:
					return value, errors.New("nie dzielimy przez 3")
				}
			}),
		}

		allRules = validator.RulesMap{
			"int":   intRules,
			"*int":  intRules,
			"**int": intRules,

			"child.id": {
				rule.Required(),
				rule.Integer[int](),
				rule.Float[float32](),
				//rule.Min(150),
			},

			"child.ancestor.id": {
				rule.Required(),
				rule.Integer[int](),
				rule.Min(500),
			},

			"array.*": {
				rule.Required(),
				rule.String(),
				rule.In([]string{"Foo", "foo"}),
			},
		}
	)

    fmt.Println("ForMapWithContext")
    forMapErrorsBag := validator.ForMapWithContext(context.Background(), mapData, allRules)
    fmt.Println(forMapErrorsBag, "\n", toJSON(forMapErrorsBag))

    fmt.Println("\nForStructWithContext")
    forStructErrorsBag := validator.ForStructWithContext(context.Background(), ptr(data{}), allRules)
    fmt.Println(forStructErrorsBag, "\n", toJSON(forStructErrorsBag))

    fmt.Println("\nForSliceWithContext")
    forSliceErrorsBag := validator.ForSliceWithContext(context.Background(), ptr([]int{1, 2, 3}),
        rule.Min(150),
    )
    fmt.Println(forSliceErrorsBag, "\n", toJSON(forSliceErrorsBag))

    fmt.Println("\nForValueWithContext (simple)")
    forValueErrors := validator.ForValueWithContext(context.Background(), ptr(ptr(6)),
        rule.Required(),
        rule.Min(150),
    )
    fmt.Println(forValueErrors, "\n", toJSON(forValueErrors))
}

type data struct {
    ID      int    `json:"id" validation:"id"`
    Name    string `json:"name" validation:"name"`
    Enabled bool   `json:"enabled" validation:"enabled"`
    Skip    string `json:"-" validation:"-"`
    NoName  string `json:"-"`
}

func ptr[T any](v T) *T {
	return &v
}

func toJSON(data any) string {
    var buf bytes.Buffer
    encoder := json.NewEncoder(&buf)
    encoder.SetIndent("", "  ")

    _ = encoder.Encode(data)

    return buf.String()
}
```

Produces:
```text
ForMapWithContext
8 field(s) failed:
int: [6][minRule{Threshold=150} minRule{Threshold=1240} minRule{Threshold=125.01} minRule{Threshold=125.01} minRule{Threshold=400} customRule{Err="nie dzielimy przez 3"}]
*int: [9][minRule{Threshold=150} minRule{Threshold=1240} minRule{Threshold=125.01} minRule{Threshold=125.01} minRule{Threshold=200} minRule{Threshold=300} maxRule{Threshold=90} maxRule{Threshold=100} minRule{Threshold=400}]
**int: [9][minRule{Threshold=150} minRule{Threshold=1240} minRule{Threshold=125.01} minRule{Threshold=125.01} minRule{Threshold=200} minRule{Threshold=300} maxRule{Threshold=90} maxRule{Threshold=100} minRule{Threshold=400}]
child.id: [2][floatRule{ExpectedType="float32", ActualType="int"} minRule{Threshold=500}]
array.1: [1][requiredRule{}]
array.2: [1][inRule{Values=[Foo foo]}]
array.3: [1][inRule{Values=[Foo foo]}]
array.4: [2][stringRule{} inRule{Values=[Foo foo]}] 
 {
  "**int": [
    {
      "rule": "MIN.NUMBER",
      "threshold": 150
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 1240
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 125.01
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 125.01
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 200
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 300
    },
    {
      "rule": "MAX.NUMBER",
      "threshold": 90
    },
    {
      "rule": "MAX.NUMBER",
      "threshold": 100
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 400
    }
  ],
  "*int": [
    {
      "rule": "MIN.NUMBER",
      "threshold": 150
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 1240
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 125.01
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 125.01
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 200
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 300
    },
    {
      "rule": "MAX.NUMBER",
      "threshold": 90
    },
    {
      "rule": "MAX.NUMBER",
      "threshold": 100
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 400
    }
  ],
  "array.1": [
    {
      "rule": "REQUIRED"
    }
  ],
  "array.2": [
    {
      "rule": "IN",
      "values": [
        "Foo",
        "foo"
      ]
    }
  ],
  "array.3": [
    {
      "rule": "IN",
      "values": [
        "Foo",
        "foo"
      ]
    }
  ],
  "array.4": [
    {
      "rule": "STRING"
    },
    {
      "rule": "IN",
      "values": [
        "Foo",
        "foo"
      ]
    }
  ],
  "child.id": [
    {
      "rule": "INT",
      "expected_type": "float32",
      "actual_type": "int"
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 500
    }
  ],
  "int": [
    {
      "rule": "MIN.NUMBER",
      "threshold": 150
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 1240
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 125.01
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 125.01
    },
    {
      "rule": "MIN.NUMBER",
      "threshold": 400
    },
    {
      "rule": "CUSTOM",
      "error": "nie dzielimy przez 3"
    }
  ]
}


ForStructWithContext
5 field(s) failed:
int: [1][requiredRule{}]
*int: [1][requiredRule{}]
**int: [1][requiredRule{}]
child.id: [2][requiredRule{} requiredRule{}]
array.*: [1][requiredRule{}] 
 {
  "**int": [
    {
      "rule": "REQUIRED"
    }
  ],
  "*int": [
    {
      "rule": "REQUIRED"
    }
  ],
  "array.*": [
    {
      "rule": "REQUIRED"
    }
  ],
  "child.id": [
    {
      "rule": "REQUIRED"
    },
    {
      "rule": "REQUIRED"
    }
  ],
  "int": [
    {
      "rule": "REQUIRED"
    }
  ]
}


ForSliceWithContext
3 field(s) failed:
0: [1][minRule{Threshold=150}]
1: [1][minRule{Threshold=150}]
2: [1][minRule{Threshold=150}] 
 {
  "0": [
    {
      "rule": "MIN.NUMBER",
      "threshold": 150
    }
  ],
  "1": [
    {
      "rule": "MIN.NUMBER",
      "threshold": 150
    }
  ],
  "2": [
    {
      "rule": "MIN.NUMBER",
      "threshold": 150
    }
  ]
}


ForValueWithContext (simple)
[minRule{Threshold=150}] 
 [
  {
    "rule": "MIN.NUMBER",
    "threshold": 150
  }
]
```
