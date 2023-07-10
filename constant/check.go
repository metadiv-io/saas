package constant

import (
	"reflect"

	"github.com/metadiv-io/env"
)

// CheckEnvSet checks if the environment variable is set.
// If not, it panics. (Only in debug or release mode)
func CheckEnvSet(e interface{}, hint ...string) {
	if env.String("GIN_MODE") == "release" || env.String("GIN_MODE") == "debug" {
		if reflect.ValueOf(e).IsZero() {
			if len(hint) > 0 {
				panic(hint[0])
			} else {
				panic("an environment variable is required")
			}
		}
	}
}
