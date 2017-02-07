package testutil

import "testing"

// AssertString that strings are equal. If not, it prints for comparison.
func AssertString(result string, expect string, t *testing.T) {
	if result != expect {
		t.Logf("\nexpect:\n%s\nresult:\n%s\n", expect, result)
		t.Fail()
	}
}
