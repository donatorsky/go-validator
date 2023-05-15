package rule

func MaxExclusive[T numberType](max T) *maxRule[T] {
	return &maxRule[T]{
		max:       max,
		inclusive: false,
	}
}
