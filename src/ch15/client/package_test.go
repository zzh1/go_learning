package client

import (
	"hello/src/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(5))
	// t.Log(series.square(5))
	t.Log(series.Square(5))
}
