package constant

import "github.com/metadiv-io/env"

var (
	MICRO_SERVICE_HOST_AUTH string
)

func init() {
	MICRO_SERVICE_HOST_AUTH = env.MustString("MICRO_SERVICE_HOST_AUTH")
}
