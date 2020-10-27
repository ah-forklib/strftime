package strftime

import (
	"testing"
)

func TestCombine(t *testing.T) {
	{
		s, err := New(`%A foo`)
		if !assertNoError(t, err, `New should succeed`) {
			return
		}

		if !assertEqual(t, 1, len(s.compiled), "there are 1 element") {
			return
		}
	}
	{
		s, _ := New(`%A 100`)
		if !assertEqual(t, 2, len(s.compiled), "there are two elements") {
			return
		}
	}
	{
		s, _ := New(`%A Mon`)
		if !assertEqual(t, 2, len(s.compiled), "there are two elements") {
			return
		}
	}
}
