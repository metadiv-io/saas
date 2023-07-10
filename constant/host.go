package constant

import "github.com/metadiv-io/env"

var (
	MICRO_SERVICE_HOST_AUTH string
)

func init() {
	MICRO_SERVICE_HOST_AUTH = env.String("MICRO_SERVICE_HOST_AUTH")
	CheckEnvSet(MICRO_SERVICE_HOST_AUTH, "MICRO_SERVICE_HOST_AUTH is required")
}
