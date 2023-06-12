package constant

import "github.com/metadiv-io/env"

var (
	MICRO_SERVICE_HOST_AUTH string
)

func init() {
	MICRO_SERVICE_HOST_AUTH = env.String("MICRO_SERVICE_HOST_AUTH")
	if MICRO_SERVICE_HOST_AUTH == "" {
		panic("MICRO_SERVICE_HOST_AUTH is required")
	}
}
