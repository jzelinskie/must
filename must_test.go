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

func ExampleGetFromMap() {
	bookmarks := map[string]string{
		"home": "https://duckduckgo.com",
	}

	fmt.Println(must.GetFromMap(bookmarks, "home"))

	// Output:
	// https://duckduckgo.com
}
