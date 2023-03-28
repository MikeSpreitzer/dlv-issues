package notime

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

func TestNoTime(t *testing.T) {
	var mg MyGood
	t.Logf("mg=%+v", mg)
	var wpw WrapPointWrap
	t.Logf("wpw=%+v", wpw)
}
