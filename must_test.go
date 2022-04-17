package must_test

import (
	"fmt"
	"net/netip"

	"github.com/jzelinskie/must"
)

func ExampleBeOk() {
	localhost := []byte{127, 0, 0, 1}
	fmt.Println(must.BeOk(netip.AddrFromSlice(localhost)))

	// Output:
	// 127.0.0.1
}
