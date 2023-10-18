package functional

// TwoCurrying is currying with two layer
func TwoCurrying[t1, t2, R any](f func(t1, t2) R) func(t1) func(t2) R {
	return func(t1 t1) func(t2) R {
		return func(t2 t2) R {
			return f(t1, t2)
		}
	}
}

// ThreeCurrying is currying with three layer
func ThreeCurrying[t1, t2, t3, R any](f func(t1, t2, t3) R) func(t1) func(t2) func(t3) R {
	return func(t1 t1) func(t2) func(t3) R {
		return func(t2 t2) func(t3) R {
			return func(t3 t3) R {
				return f(t1, t2, t3)
			}
		}
	}
}
