package main

import (
	"fmt"
)

type numbersSlice []int

func newNumbersSlice(length int) numbersSlice {

	ns := numbersSlice{}

	for a := 1; a <= length; a++ {
		ns = append(ns, a)
	}

	return ns
}

// func (ns numbersSlice) addNewNumber(number int) numbersSlice {
// 	ns = append(ns, number)

// 	return ns
// }

func checkSliceNumbers(slice numbersSlice) {

	for _, n := range slice {
		if n%2 > 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}

}
