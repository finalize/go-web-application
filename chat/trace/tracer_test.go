package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return value is nil from New")
	} else {
		tracer.Trace("Hello")
		if buf.String() != "Hello\n" {
			t.Errorf("'%s'という誤った文字列が検出されました", buf.String())
		}
	}
	t.Error("Not Creating test yet")
}
