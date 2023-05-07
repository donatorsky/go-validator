package rule

func BetweenExclusive[T numberType](min, max T) *betweenRule[T] {
	return &betweenRule[T]{
		min:       min,
		max:       max,
		inclusive: false,
	}
}
