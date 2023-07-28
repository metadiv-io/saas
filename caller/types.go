package caller

import (
	"github.com/metadiv-io/ginger/caller"
	"github.com/metadiv-io/saas/types"
)

type Response[T any] struct {
	caller.Response[T]
	Credit float64       `json:"credit"`
	Traces []types.Trace `json:"traces,omitempty"`
}
