package hid_test

import (
	"fmt"
	"testing"

	"github.com/dh1tw/hid"
)

// Tests that device enumeration can be called concurrently from multiple threads.
func TestEnumerate(t *testing.T) {
	list := hid.Enumerate(4057, 96)
	for _, l := range list {
		fmt.Printf("%#v\n", l)
	}
}
