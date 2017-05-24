package eid_test

import (
	"testing"
	"github.com/jchavannes/iiproject/eid"
)

const (
	TestEidShort = "a@myeid.org"
	TestEidFull = "myeid.org/id/a"
)

func TestUrl(t *testing.T) {
	full := eid.ConvertShortEidUrlIntoFull(TestEidShort)
	if full != TestEidFull {
		t.Errorf("Error, full does not match full eid - expecting: %s, got: %s", TestEidFull, full)
	}

	short := eid.ConvertFullEidUrlIntoShort(TestEidFull)
	if short != TestEidShort{
		t.Errorf("Error, short does not match short eid- expecting: %s, got: %s", TestEidShort, short)
	}
}
