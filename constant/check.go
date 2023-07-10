package constant

import (
	"reflect"

	"github.com/metadiv-io/env"
)

// CheckEnvSet checks if the environment variable is set.
// If not, it panics. (Skip if GIN_MODE == test)
func CheckEnvSet(e interface{}) {
	if env.String("GIN_MODE") == "test" {
		return
	}
	if reflect.ValueOf(e).IsZero() {
		panic(reflect.TypeOf(e).Name() + " is required")
	}
}
