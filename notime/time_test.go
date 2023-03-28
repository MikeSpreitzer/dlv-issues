package main

import (
	"testing"
	"time"
)

type MyGood struct {
	TheTime *time.Time
}

type WrapPointWrap struct {
	Point *Time
}

func TestGood(t *testing.T) {
	var mg MyGood
	t.Logf("mg=%+v", mg)
}

func TestBad(t *testing.T) {
	var wpw WrapPointWrap
	t.Logf("wpw=%+v", wpw)
}
