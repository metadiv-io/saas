package types

import (
	"github.com/metadiv-io/ginger/types"
)

type Response struct {
	types.Response
	Traces []Trace `json:"traces"`
	Credit float64 `json:"credit"`
}

func (r *Response) Calculate() {
	r.Credit = 0
	r.Duration = 0
	if r.Traces != nil {
		for _, trace := range r.Traces {
			r.Duration += trace.Duration
			r.Credit += trace.Credit
		}
	}
}
