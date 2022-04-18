// Must implements utility functions that panic when an assumption isn't met.
//
// This is meant to reduce boiler plate. Please panic responsibly.
package must

// BeOk panics if ok is false.
func BeOk[T any](v T, ok bool) T {
	if !ok {
		panic("must have been ok")
	}
	return v
}

// NotBeOk panics if ok is true.
func NotBeOk[T any](v T, ok bool) T {
	if ok {
		panic("must have been ok")
	}
	return v
}

// MustGetFromMap panics if the key is not in the provided map.
func GetFromMap[T comparable, Y any](gmap map[T]Y, key T) Y {
	v, ok := gmap[key]
	return BeOk(v, ok)
}

// BeOne panics if the slice does not have only one element.
func BeOne[T any](v []T) T {
	if len(v) != 1 {
		panic("must be slice with only one element")
	}
	return v[0]
}

// NotError panics if err is not equal to nil.
func NotError[T any](v T, err error) T {
	if err != nil {
		panic("must have not fail: " + err.Error())
	}
	return v
}
