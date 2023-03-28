package notime

import (
	"testing"
	"time"
)

type MyBad struct {
	TheTime *time.Time
}

func TestNoTime(t *testing.T) {
	var mb MyBad
	t.Logf("mb=%+v", mb)
}
