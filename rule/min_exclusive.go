package rule

func MinExclusive[T numberType](min T) *minRule[T] {
	return &minRule[T]{
		min:       min,
		inclusive: false,
	}
}
