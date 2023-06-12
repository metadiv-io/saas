package types

import "github.com/metadiv-io/sql"

type Response struct {
	Success    bool            `json:"success"`
	TraceID    string          `json:"trace_id"`
	Locale     string          `json:"locale"`
	Duration   int64           `json:"duration"`
	Credit     float64         `json:"credit"`
	Pagination *sql.Pagination `json:"pagination,omitempty"`
	Error      *Error          `json:"error,omitempty"`
	Data       interface{}     `json:"data,omitempty"`
	Traces     []Trace         `json:"traces,omitempty"`
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
