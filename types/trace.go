package types

import (
	"github.com/metadiv-io/ginger/types"
)

type Trace struct {
	types.Trace
	Credit float64 `json:"credit"`
}
