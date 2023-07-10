package constant

import (
	"reflect"

	"github.com/metadiv-io/env"
)

// CheckEnvSet checks if the environment variable is set.
// If not, it panics. (Skip if GIN_MODE == test)
func CheckEnvSet(e interface{}, hint ...string) {
	if env.String("GIN_MODE") == "test" {
		return
	}
	if reflect.ValueOf(e).IsZero() {
		if len(hint) > 0 {
			panic(hint[0])
		} else {
			panic("an environment variable is required")
		}
	}
}
