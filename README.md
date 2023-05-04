# Go Validator
A Laravel-like data validator for Go.

[![GitHub license](https://img.shields.io/github/license/donatorsky/go-validator)](https://github.com/donatorsky/go-validator/blob/main/LICENSE)
[![Build](https://github.com/donatorsky/go-validator/workflows/Tests/badge.svg?branch=main)](https://github.com/donatorsky/go-validator/actions?query=branch%3Amain)

## Installation

```shell
go get github.com/donatorsky/go-validator
```

The library is still a work in progress. More details soon.

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

		allRules = RulesMap{
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
