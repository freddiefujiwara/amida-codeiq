package main

import (
	"reflect"
	"testing"
)

func TestListUp(t *testing.T) {
	array := []int{1, 2, 3, 4}
	amida := []int{1, 0, 0}
	actual := Amida(array, amida)
	expected := []int{2, 1, 3, 4}
	if !reflect.DeepEqual(len(actual), len(expected)) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	for _, v := range [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}, {0, 0, 1}} {
		actual = Amida(array, v)
	}
	expected = []int{3, 1, 2, 4}
	if !reflect.DeepEqual(len(actual), len(expected)) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
