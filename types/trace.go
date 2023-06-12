package types

type Trace struct {
	Success    bool    `json:"success"`
	SystemUUID string  `json:"system_uuid"`
	SystemName string  `json:"system_name"`
	TraceID    string  `json:"trace_id"`
	Time       int64   `json:"time"`
	Duration   int64   `json:"duration"`
	Credit     float64 `json:"credit"`
	Error      *Error  `json:"error,omitempty"`
}
