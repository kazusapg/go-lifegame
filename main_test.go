package main

import (
	"reflect"
	"testing"
)

func TestNext(t *testing.T) {
	gridLength := 10
	grid := makeTestGrid(t)
	gotGrid := deepCopySlice(grid, gridLength)
	next(grid, gotGrid, gridLength)
	wantGrid := makeWantGrid(t)

	if !reflect.DeepEqual(gotGrid, wantGrid) {
		t.Errorf("got %v want %v", gotGrid, wantGrid)
	}
}

func makeTestGrid(t *testing.T) Grid {
	t.Helper()
	l := life
	n := nolife
	testGrid := Grid{
		[]string{n, n, n, n, n, l, n, l, l, l},
		[]string{l, n, l, n, l, l, n, n, n, l},
		[]string{n, l, l, l, n, n, n, l, l, n},
		[]string{n, l, n, n, n, l, n, n, n, l},
		[]string{n, l, l, n, n, n, l, l, n, n},
		[]string{n, n, n, n, l, n, l, l, n, n},
		[]string{n, l, n, l, l, l, n, l, n, l},
		[]string{n, l, l, l, n, l, n, n, l, l},
		[]string{n, n, n, n, n, n, n, n, n, l},
		[]string{l, l, l, n, l, n, l, n, n, l},
	}
	return testGrid
}

func makeWantGrid(t *testing.T) Grid {
	t.Helper()
	l := life
	n := nolife
	testGrid := Grid{
		[]string{n, n, n, n, l, l, l, n, l, l},
		[]string{n, n, l, n, l, l, n, n, n, l},
		[]string{l, n, n, l, n, l, l, n, l, l},
		[]string{l, n, n, l, n, n, n, n, n, n},
		[]string{n, l, l, n, n, n, n, l, l, n},
		[]string{n, l, n, n, l, n, n, n, n, n},
		[]string{n, l, n, n, n, n, n, l, n, l},
		[]string{n, l, n, l, n, l, l, n, n, l},
		[]string{l, n, n, n, l, l, n, n, n, l},
		[]string{n, l, n, n, n, n, n, n, n, n},
	}
	return testGrid
}
