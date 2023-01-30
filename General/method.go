package General

import "net/http"

// Method
const (
	MethodGet     = http.MethodGet
	MethodHead    = http.MethodHead
	MethodPost    = http.MethodPost
	MethodPut     = http.MethodPut
	MethodPatch   = http.MethodPatch
	MethodDelete  = http.MethodDelete
	MethodConnect = http.MethodConnect
	MethodOptions = http.MethodOptions
	MethodTrace   = http.MethodTrace
)

// ContentType
const (
	ContentJson = "application/json"
	ContentText = "text/plain"
	ContentHtml = "text/html"
)

