package main

import (
	"fmt"
	"reflect"
)

func main() {
	m1 := map[int]interface{}{}
	m2 := map[int]interface{}{}
	count := 10000
	for i := 0; i < count; i++ {
		m1[i] = i
		m2[count-i-1] = count - i - 1
	}
	fmt.Println(reflect.DeepEqual(m1, m2))
}
