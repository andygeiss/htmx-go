package testdata_test

import (
	"reflect"
	"testing"
)

func Assert(t *testing.T, desc string, got any, expected any) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("%s, but got %v and expected %v", desc, got, expected)
	}
}
