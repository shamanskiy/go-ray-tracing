package core

func IfElse[T any](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	} else {
		return ifFalse
	}
}
