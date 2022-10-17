package util

import (
	"fmt"
)

type MyInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func PrintMyInteger[T MyInteger](i T) {
	fmt.Printf("Value is %d\n", i)
}
