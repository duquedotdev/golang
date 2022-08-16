package main

import "testing"

func TestNewSliceOfNumbers(t *testing.T) {
	newSlice := newNumbersSlice(10)

	if newSlice == nil {
		t.Errorf("Slice is not instantiated")
	}

	if len(newSlice) != 10 {
		t.Errorf("Slice is incomplete %v", len(newSlice))
	}
}
