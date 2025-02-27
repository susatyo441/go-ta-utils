package functions

import "reflect"

// IsZeroValue checks if a given struct is equal to its zero value.
// The zero value of a struct is the state where all its fields are uninitialized:
// - Numeric fields are 0.
// - String fields are "".
// - Boolean fields are false.
// - Pointers, slices, maps, and interfaces are nil.
// This function is particularly useful for structs, where field-by-field comparison
// would be tedious and error-prone.
//
// Parameters:
// - model: An interface{} representing the struct to check.
//
// Returns:
// - bool: true if the struct is equal to its zero value, false otherwise.
//
// Usage: This function simplifies checking if a struct is "empty" or in its default state,
// which can be difficult to do manually. It's especially helpful when dealing with nested structs
// or dynamic types where the struct's fields may not be known at compile time.
func IsZeroValue(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}
