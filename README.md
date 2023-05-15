# Go Validator

A Laravel-like data validator for Go.

[![GitHub license](https://img.shields.io/github/license/donatorsky/go-validator)](https://github.com/donatorsky/go-validator/blob/main/LICENSE)
[![Build](https://github.com/donatorsky/go-validator/workflows/Tests/badge.svg?branch=main)](https://github.com/donatorsky/go-validator/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/donatorsky/go-validator/branch/main/graph/badge.svg?token=LNSP9QBPQS)](https://codecov.io/gh/donatorsky/go-validator)

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

It is possible to validate nested objects (i.e.: slice, array, map or struct) using the dot notation:

- For slices and arrays: it refers to the index of element, e.g.: given `"slice": []int{1, 2, 3},`, `slice.1` refers to the value `2`.
- For maps: it refers to the element by given key, e.g.: given `"map": map[string]int{"foo": 1, "bar": 2, "baz": 3},`, `map.bar` refers to the value `2`.
- For structs: it refers to the field with same `validation` tag or field name if tag is not present, e.g.: given `"struct": {Foo: 1, Bar: 2, Baz: 3},`, `struct.Bar` refers to the value `2`.

You can also validate every single value of slice and array by using `*` wildcard symbol.

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
        rule.WhenFunc(
            func(ctx context.Context, value any, data any) bool {
                value, isNil := rule.Dereference(value)
				if isNil {
					return false
                }
				
                return value.(int)%2 == 1
            },
            rule.Min(100),
            rule.WhenFunc(
                func(_ context.Context, value any, _ any) bool {
                    value, isNil := rule.Dereference(value)
                    if isNil {
                        return false
                    }
                    
                    return value.(int)%2 == 0
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

Common types:

- `integerType` - Any integer type, i.e.: `~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64`.
- `floatType` - Any integer type, i.e.: `~float32 | ~float64`.
- `numberType` - Any number, i.e.: `integerType | floatType`.
- `afterComparable` - Any object that implements the following interface:
```go
type afterComparable interface {
	After(time.Time) bool
}
```
- `afterOrEqualComparable` - Any object that implements the following interface:
```go
type afterOrEqualComparable interface {
    Equal(time.Time) bool
    After(time.Time) bool
}
```
- `beforeComparable` - Any object that implements the following interface:
```go
type beforeComparable interface {
	Before(time.Time) bool
}
```
- `beforeOrEqualComparable` - Any object that implements the following interface:
```go
type beforeOrEqualComparable interface {
    Equal(time.Time) bool
    Before(time.Time) bool
}
```
- `Comparator` - Any object of the following type:
```go
type Comparator func(x, y any) bool
```

### `After(after time.Time)`

Checks whether a value is after `after` date.

**Applies to:**

Any value. Passes only when a value implements `afterComparable` interface and is after `after` date.

**Modifies output:**

No.

**Bails:**

No.

### `AfterOrEqual(afterOrEqual time.Time)`

Checks whether a value is after or equal to `afterOrEqual` date.

**Applies to:**

Any value. Passes only when a value implements `afterOrEqualComparable` interface and is after or equal to `afterOrEqual` date.

**Modifies output:**

No.

**Bails:**

No.

### `Array()`

Checks and ensures that a value is of array type.

**Applies to:**

Any value. Passes only when a value is of array type or its pointer, any length and element type.

**Modifies output:**

No.

**Bails:**

Yes, when a value is not of array type.

### `ArrayOf[Out any]()`

Checks and ensures that a value is of `[n]Out` type.

**Applies to:**

Any value. Passes only when a value is of `[n]Out` type or its pointer, any length, or `nil` array of any type.

**Modifies output:**

Yes. Returns `nil` slice of `[0]Out` type for `nil` values. Returns input value otherwise.

**Bails:**

Yes, when a value is not of `[n]Out` type or its pointer.

### `Before(before time.Time)`

Checks whether a value is before `before` date.

**Applies to:**

Any value. Passes only when a value implements `beforeComparable` interface and is before `before` date.

**Modifies output:**

No.

**Bails:**

No.

### `BeforeOrEqual(beforeOrEqual time.Time)`

Checks whether a value is before or equal to `beforeOrEqual` date.

**Applies to:**

Any value. Passes only when a value implements `beforeOrEqualComparable` interface and is before or equal to `beforeOrEqual` date.

**Modifies output:**

No.

**Bails:**

No.

### `Between[T numberType](min, max T)`

Checks whether a value is between `min` and `max`, inclusive.

**Applies to:**

- `numberType`: checks if a value is between `min` and `max`.
- `string`: checks if string's length is between `min` and `max`.
- `slice`, `array`: checks if slice/array has at least `min` and `max` elements.
- `map`: checks if map has at least `min` and `max` keys.

**Modifies output:**

No.

**Bails:**

No.

### `BetweenExclusive[T numberType](min, max T)`

Checks whether a value is between `min` and `max`, exclusive.

**Applies to:**

- `numberType`: checks if a value is between `min` and `max`.
- `string`: checks if string's length is between `min` and `max`.
- `slice`, `array`: checks if slice/array has at least `min` and `max` elements.
- `map`: checks if map has at least `min` and `max` keys.

**Modifies output:**

No.

**Bails:**

No.

### `Boolean()`

Checks and ensures that a value is of `bool` type.

**Applies to:**

- `bool`: any value.
- `integerType`: when a value equals to `0` or `1`.
- `floatType`: when a value equals to `0.0` or `1.0`.
- `string`: when string is convertible to `bool` according to the `strconv.ParseBool` function.

**Modifies output:**

- `nil`: returns `*bool` nil pinter.
- `bool`: returns unchanged value.
- `integerType`: `false` when `0`, `true` when `1`.
- `floatType`: `false` when `0.0`, `true` when `1.0`.
- `string`: according to the `strconv.ParseBool` function.

**Bails:**

Yes, when a value is not of `bool` type, its pointer or cannot be converted to it.

### `Date()`

Checks whether a value is of `time.Time` type, its pointer or valid date string in `time.RFC3339Nano` format.

**Applies to:**

- `nil`.
- `time.Duration`: any value.
- `string`: when string is convertible to `time.Duration` according to the `time.Parse` function and `time.RFC3339Nano` format.

**Modifies output:**

- `nil`: returns `*time.Time` nil pinter.
- `time.Time`: returns unchanged value.
- `string`: according to the `time.Parse` function.

**Bails:**

No.

### `DateFormat(format string)`

Checks whether a value is of `time.Time` type, its pointer or valid date string in `format` format.

**Applies to:**

- `nil`.
- `time.Duration`: any value.
- `string`: when string is convertible to `time.Duration` according to the `time.Parse` function and `format` format.

**Modifies output:**

- `nil`: returns `*time.Time` nil pinter.
- `time.Time`: returns unchanged value.
- `string`: according to the `time.Parse` function.

**Bails:**

No.

### `DoesntEndWith(suffix string, suffixes ...string)`

Checks whether a value is a string not ending with any of provided suffixes.

**Applies to:**

- `nil`: passes.
- `string`: checks if string does not end with any of provided suffixes (case-sensitive).
- `any`: fails.

**Modifies output:**

No.

**Bails:**

No.

### `DoesntStartWith(prefix string, prefixes ...string)`

Checks whether a value is a string not starting with any of provided prefixes.

**Applies to:**

- `nil`: passes.
- `string`: checks if string does not start with any of provided prefixes (case-sensitive).
- `any`: fails.

**Modifies output:**

No.

**Bails:**

No.

### `Duration()`

Checks whether a value is of `time.Duration` type, its pointer or valid duration string.

**Applies to:**

- `time.Duration`: any value.
- `string`: when string is convertible to `time.Duration` according to the `time.ParseDuration` function.

**Modifies output:**

- `nil`: returns `*time.Duration` nil pinter.
- `time.Duration`: returns unchanged value.
- `string`: according to the `time.ParseDuration` function.

**Bails:**

No.

### `Email()`

Checks whether a value is valid email address, according to the `net/mail.ParseAddress` function.

**Applies to:**

Any value. Passes only when a value is of `string` type or its pointer and was successfully parsed by `net/mail.ParseAddress` function.

**Modifies output:**

No.

**Bails:**

No.

### `EmailAddress()`

Checks whether a value is valid email address, according to the `net/mail.ParseAddress` function.

**Applies to:**

Any value. Passes only when a value is of `string` type or its pointer and was successfully parsed by `net/mail.ParseAddress` function.

**Modifies output:**

Yes. Unlike the `Email()` rule, it returns the email address of the string. E.g. given a value `Foo Bar <foo@bar.baz> (some comment)`, the output will be `foo@bar.baz`.

**Bails:**

No.

### `EndsWith(suffix string)`

Checks whether a value is a string ending with one of provided suffixes.

**Applies to:**

- `nil`: passes.
- `string`: checks if string ends with one of provided suffixes (case-sensitive).
- `any`: fails.

**Modifies output:**

No.

**Bails:**

No.

### `Filled()`

Checks whether a value is not empty when it is present.

**Applies to:**

- `nil`: passes.
- `*any`: passes.
- `!nil`: checks if value is not a zero value.

**Modifies output:**

No.

**Bails:**

No.

### `Float[Out floatType]()`

Checks and ensures that a value is of `Out` type or its pointer.

**Applies to:**

Any value. Passes only when a value is of `Out` type or its pointer.

**Modifies output:**

No.

**Bails:**

Yes, when a value is not of `Out` type or its pointer.

### `In[T comparable](values []T, options ...inRuleOption)`

Checks whether a value exists in `values`.

**Options:**

- `InRuleWithComparator(comparator Comparator)`: sets custom elements comparator. `comparator` receives an input value and each element of `values`, one at a time.
- `InRuleWithoutAutoDereference()`: disables automatic dereference of a value, i.e. `values` will be compared against the exact input value which may be a pointer.

**Applies to:**

Any value. Passes only when a value exists in `values`, optionally using custom `comparator`.
When auto dereference is enabled (and it is by default), `nil` value will pass validation.

**Modifies output:**

No.

**Bails:**

No.

### `Integer[Out integerType]()`

Checks and ensures that a value is of `Out` type or its pointer.

**Applies to:**

Any value. Passes only when a value is of `Out` type or its pointer.

**Modifies output:**

No.

**Bails:**

Yes, when a value is not of `Out` type or its pointer.

### `IP()`

Checks whether a value is a string in IP v4 or v6 format.

**Applies to:**

- `nil`: passes.
- `string`: checks if a value is in IP v4 or v6 format according to the `net.ParseIP` function.

**Modifies output:**

No.

**Bails:**

No.

### `Length[T integerType](length T)`

Checks whether a value is exactly `length`.

**Applies to:**

- `string`: checks if string's length is exactly `length` characters.
- `slice`, `array`: checks if slice/array has exactly `length` elements.
- `map`: checks if map has exactly `length` keys.

**Modifies output:**

No.

**Bails:**

No.

### `Map()`

Checks and ensures that a value is of map type.

**Applies to:**

Any value. Passes only when a value is of map type or its pointer, any length and key/value type.

**Modifies output:**

No.

**Bails:**

Yes, when a value is not of map type.

### `Max[T numberType](max T)`

Checks whether a value is at most `max`.

**Applies to:**

- `nil`: passes.
- `numberType`: checks if a value is at most `max` .
- `string`: checks if string's length is at most `max` characters.
- `slice`, `array`: checks if slice/array has at most `max` elements.
- `map`: checks if map has at most `max` keys.

**Modifies output:**

No.

**Bails:**

No.

### `MaxExclusive[T numberType](max T)`

Checks whether a value is less than `max`.

**Applies to:**

- `nil`: passes.
- `numberType`: checks if a value is less than `max` .
- `string`: checks if string's length is less than `max` characters.
- `slice`, `array`: checks if slice/array has less than `max` elements.
- `map`: checks if map has less than `max` keys.

**Modifies output:**

No.

**Bails:**

No.

### `Min[T numberType](min T)`

Checks whether a value is at least `min`.

**Applies to:**

- `nil`: passes.
- `numberType`: checks if a value is at least `min`.
- `string`: checks if string's length is at least `min` characters.
- `slice`, `array`: checks if slice/array has at least `min` elements.
- `map`: checks if map has at least `min` keys.

**Modifies output:**

No.

**Bails:**

No.

### `MinExclusive[T numberType](min T)`

Checks whether a value is greater than `min`.

**Applies to:**

- `nil`: passes.
- `numberType`: checks if a value is greater than `min`.
- `string`: checks if string's length is more than `min` characters.
- `slice`, `array`: checks if slice/array has more than `min` elements.
- `map`: checks if map has more than `min` keys.

**Modifies output:**

No.

**Bails:**

No.

### `NotIn[T comparable](values []T, options ...notInRuleOption)`

Checks whether a value does not exist in `values`.

**Options:**

- `InRuleWithComparator(comparator Comparator)`: sets custom elements comparator. `comparator` receives an input value and each element of `values`, one at a time.
- `InRuleWithoutAutoDereference()`: disables automatic dereference of a value, i.e. `values` will be compared against the exact input value which may be a pointer.

**Applies to:**

Any value. Passes only when a value does not exist in `values`, optionally using custom `comparator`.
When auto dereference is enabled (and it is by default), `nil` value will pass validation.

**Modifies output:**

No.

**Bails:**

No.

### `NotRegex(regex *regexp.Regexp)`

Checks whether a value does not match `regex` expression.

**Applies to:**

- `nil`: passes.
- `string`: checks if string does not match `regex` expression.
- `any`: fails.

**Modifies output:**

No.

**Bails:**

No.

### `Numeric()`

Checks whether a value is a numeric value or a string that can be converted to one.

**Applies to:**

- `nil`: passes.
- `numberType`, `complex64`, `complex128`: passes.
- `string`: passes only when a value can be converted to number using `strconv.ParseInt`, `strconv.ParseUint`, `strconv.ParseFloat` and `strconv.ParseComplex`, in that order.
- `any`: fails.

**Modifies output:**

- `nil`: unchanged value.
- `numberType`, `complex64`, `complex128`: unchanged value.
- `string`: value converted to number.

**Bails:**

No.

### `Regex(regex *regexp.Regexp)`

Checks whether a value matches `regex` expression.

**Applies to:**

- `nil`: passes.
- `string`: checks if string matches `regex` expression.
- `any`: fails.

**Modifies output:**

No.

**Bails:**

No.

### `Required()`

Checks whether a value is not nil.

**Applies to:**

Any value.

**Modifies output:**

No.

**Bails:**

Yes. When value is nil.

### `Slice()`

Checks and ensures that a value is of slice type.

**Applies to:**

Any value. Passes only when a value is of slice type or its pointer, any length and element type.

**Modifies output:**

No.

**Bails:**

Yes, when a value is not of slice type or its pointer.

### `SliceOf[Out any]()`

Checks and ensures that a value is of `[]Out` type.

**Applies to:**

Any value. Passes only when a value is of `[]Out` type or its pointer, any length, or `nil` slice of any type.

**Modifies output:**

Yes. Returns `nil` slice of `[]Out` type for `nil` values. Returns input value otherwise.

**Bails:**

Yes, when a value is not of `[]Out` type or its pointer.

### `StartsWith(prefix string, prefixes ...string)`

Checks whether a value is a string starting with one of provided prefixes.

**Applies to:**

- `nil`: passes.
- `string`: checks if string starts with one of provided prefixes (case-sensitive).
- `any`: fails.

**Modifies output:**

No.

**Bails:**

No.

### `String()`

Checks and ensures that a value is of `string` type or its pointer.

**Applies to:**

Any value. Passes only when a value is of `string` type or its pointer.

**Modifies output:**

No.

**Bails:**

Yes, when a value is not of `string` type or its pointer.

### `Struct()`

Checks and ensures that a value is of struct type.

**Applies to:**

Any value. Passes only when a value is of struct type or its pointer, any type.

**Modifies output:**

No.

**Bails:**

Yes, when a value is not of struct type.

### `URL()`

Checks whether a value is a valid URL string.

**Applies to:**

- `nil`: passes.
- `string`: checks if string is a valid URL according to `net/url.ParseRequestURI` function.
- `any`: fails.

**Modifies output:**

No.

**Bails:**

No.

### `UUID(options ...uuidRuleOption)`

Checks whether a value is a valid RFC 4122 (version 1, 3, 4 or 5) universally unique identifier (UUID).

**Options:**

- `UUIDRuleVersion1()`: allows for UUIDv1.
- `UUIDRuleVersion3()`: allows for UUIDv3.
- `UUIDRuleVersion4()`: allows for UUIDv4.
- `UUIDRuleVersion5()`: allows for UUIDv5.
- `UUIDRuleDisallowNilUUID()`: disallows for nil UUID, i.e. `00000000-0000-0000-0000-000000000000`.

**Applies to:**

- `nil`: passes.
- `string`: checks if string is a valid UUID.
- `any`: fails.

**Modifies output:**

No.

**Bails:**

No.

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

## JSON formatting

Both `error.ValidationError` and `error.ErrorsBag` support JSON marshalling giving your application a handy way of reporting errors that occured.

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/donatorsky/go-validator"
    "github.com/donatorsky/go-validator/rule"
)

func main() {
    errorsBag := validator.ForMap(map[string]any{...}, validator.RulesMap{...})
    fmt.Println(errorsBag)
    fmt.Println()
    printJSON(errorsBag)
}

func printJSON(data any) {
    encoder := json.NewEncoder(os.Stdout)
    encoder.SetIndent("", "  ")

    _ = encoder.Encode(data)
}
```

Produces something similar to:

```text
4 field(s) failed:
int: [1][must be at least 150]
child.id: [1][must be an int but is float64]
child.roles.*: [1][is required]
array.4: [2][must end with "oo"; does not exist in [Foo foo]}]

{
  "int": [
    {
      "rule": "MIN.NUMBER"
      "threshold": 150
    }
  ],
  "child.id": [
    {
      "rule": "INT"
      "expected_type": "int"
      "actual_type": "float64"
    }
  ],
  "child.roles.*": [
    {
      "rule": "REQUIRED"
    }
  ],
  "array.4": [
    {
      "rule": "ENDS_WITH"
      "end_part": "oo"
    },
    {
      "rule": "IN",
      "values": [
        "Foo",
        "foo"
      ]
    }
  ],
}
```
